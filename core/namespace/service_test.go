package namespace_test

import (
	"context"
	testing "testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/goto/siren/core/namespace"
	"github.com/goto/siren/core/namespace/mocks"
	"github.com/goto/siren/core/provider"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/secret"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

const testProviderType = "test-type"

func TestService_ListNamespaces(t *testing.T) {

	type testCase struct {
		Description        string
		ExpectedNamespaces []namespace.Namespace
		Setup              func(*mocks.Repository, *mocks.Encryptor, testCase)
		Err                error
	}
	var (
		ctx       = context.TODO()
		timeNow   = time.Now()
		testCases = []testCase{
			{
				Description: "should return error if List repository error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List(mock.AnythingOfType("context.todoCtx")).Return(nil, errors.New("some error"))
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if List repository success and decrypt error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List(mock.AnythingOfType("context.todoCtx")).Return([]namespace.EncryptedNamespace{
						{
							Namespace: &namespace.Namespace{
								ID: 1,
								Provider: provider.Provider{
									ID: 1,
								},
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							CredentialString: `encrypted-text-1`,
						},
						{
							Namespace: &namespace.Namespace{
								ID: 2,
								Provider: provider.Provider{
									ID: 1,
								},
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							CredentialString: `encrypted-text-2`,
						},
					}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("decrypt error"))
				},
				Err: errors.New("decrypt error"),
			},
			{
				Description: "should return error if list repository success and decrypted object is not json",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List(mock.AnythingOfType("context.todoCtx")).Return([]namespace.EncryptedNamespace{
						{
							Namespace: &namespace.Namespace{
								ID: 1,
								Provider: provider.Provider{
									ID: 1,
								},
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							CredentialString: `encrypted-text-1`,
						},
						{
							Namespace: &namespace.Namespace{
								ID: 2,
								Provider: provider.Provider{
									ID: 1,
								},
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							CredentialString: `encrypted-text-2`,
						},
					}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("", nil)
				},
				Err: errors.New("unexpected end of JSON input"),
			},
			{
				Description: "should success if list repository and decrypt success",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List(mock.AnythingOfType("context.todoCtx")).Return([]namespace.EncryptedNamespace{
						{
							Namespace: &namespace.Namespace{
								ID: 1,
								Provider: provider.Provider{
									ID: 1,
								},
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							CredentialString: `encrypted-text-1`,
						},
						{
							Namespace: &namespace.Namespace{
								ID: 2,
								Provider: provider.Provider{
									ID: 1,
								},
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							CredentialString: `encrypted-text-2`,
						},
					}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{\"name\": \"a\"}", nil)
				},
				ExpectedNamespaces: []namespace.Namespace{
					{
						ID: 1,
						Provider: provider.Provider{
							ID: 1,
						},
						Name:   "foo",
						Labels: map[string]string{"foo": "bar"},
						Credentials: map[string]any{
							"name": "a",
						},
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					},
					{
						ID: 2,
						Provider: provider.Provider{
							ID: 1,
						},
						Name:   "foo",
						Labels: map[string]string{"foo": "bar"},
						Credentials: map[string]any{
							"name": "a",
						},
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					},
				},
				Err: nil,
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock = new(mocks.Repository)
				encryptorMock  = new(mocks.Encryptor)
			)
			svc := namespace.NewService(encryptorMock, repositoryMock, nil, nil)

			tc.Setup(repositoryMock, encryptorMock, tc)

			got, err := svc.List(ctx)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}
			if !cmp.Equal(got, tc.ExpectedNamespaces) {
				t.Fatalf("got result %+v, expected was %+v", got, tc.ExpectedNamespaces)
			}
			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestService_CreateNamespace(t *testing.T) {
	type testCase struct {
		Description string
		NSpace      *namespace.Namespace
		Setup       func(*mocks.Repository, *mocks.Encryptor, *mocks.ProviderService, *mocks.ConfigSyncer, testCase)
		Err         error
	}
	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "should return error if namespace is nil",
				Err:         errors.New("incoming namespace is empty"),
			},
			{
				Description: "should return error if provider service return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(nil, errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if encrypt return error caused credential is not in json",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"invalid": make(chan int),
					},
				},
				Err: errors.New("json: unsupported type: chan int"),
			},
			{
				Description: "should return error if encrypt return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if create repository error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error conflict if create repository return duplicate error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(namespace.ErrDuplicate)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("urn and provider pair already exist"),
			},
			{
				Description: "should return error not found if create repository return relation error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(namespace.ErrRelation)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("provider id does not exist"),
			},
			{
				Description: "should return error if create repository success & sync config return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(nil)
					cs.EXPECT().SyncRuntimeConfig(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("provider.Provider")).Return(nil, errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if create repository success & sync config success & update labels error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(nil)
					cs.EXPECT().SyncRuntimeConfig(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("provider.Provider")).Return(map[string]string{"k": "v"}, nil)
					rr.EXPECT().UpdateLabels(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return nil error if create repository success & sync config success & update labels success",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, ps *mocks.ProviderService, cs *mocks.ConfigSyncer, tc testCase) {
					ps.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&provider.Provider{Type: testProviderType}, nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(nil)
					cs.EXPECT().SyncRuntimeConfig(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("provider.Provider")).Return(map[string]string{"k": "v"}, nil)
					rr.EXPECT().UpdateLabels(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: nil,
			},
		}
	)
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock      = new(mocks.Repository)
				encryptorMock       = new(mocks.Encryptor)
				providerServiceMock = new(mocks.ProviderService)
				providerPluginMock  = new(mocks.ConfigSyncer)
			)
			svc := namespace.NewService(encryptorMock, repositoryMock, providerServiceMock,
				map[string]namespace.ConfigSyncer{
					testProviderType: providerPluginMock,
				},
			)

			if tc.Setup != nil {
				tc.Setup(repositoryMock, encryptorMock, providerServiceMock, providerPluginMock, tc)
			}

			err := svc.Create(ctx, tc.NSpace)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}

			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestService_GetNamespace(t *testing.T) {
	type testCase struct {
		Description string
		NSpace      *namespace.Namespace
		Setup       func(*mocks.Repository, *mocks.Encryptor, testCase)
		Err         error
	}
	var (
		ctx       = context.TODO()
		testID    = uint64(10)
		testCases = []testCase{
			{
				Description: "should return error if Get repository error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), testID).Return(nil, errors.New("some error"))
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error not found if Get repository return not found error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), testID).Return(nil, namespace.NotFoundError{})
				},
				Err: errors.New("namespace not found"),
			},
			{
				Description: "should return error if Get repository success and decrypt return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), testID).Return(&namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}, nil)
					e.EXPECT().Decrypt(secret.MaskableString("some-ciphertext")).Return("", errors.New("some error"))
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if Get repository success and decrypted credentials is not json marshallable",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), testID).Return(&namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}, nil)
					e.EXPECT().Decrypt(secret.MaskableString("some-ciphertext")).Return("", nil)
				},
				Err: errors.New("unexpected end of JSON input"),
			},
			{
				Description: "should return nil error if Get repository success and decrypt success",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), testID).Return(&namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}, nil)
					e.EXPECT().Decrypt(secret.MaskableString("some-ciphertext")).Return("{ \"key\": \"value\" }", nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"key": "value",
					},
				},
				Err: nil,
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock = new(mocks.Repository)
				encryptorMock  = new(mocks.Encryptor)
			)
			svc := namespace.NewService(encryptorMock, repositoryMock, nil, nil)

			tc.Setup(repositoryMock, encryptorMock, tc)

			got, err := svc.Get(ctx, testID)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}
			if !cmp.Equal(got, tc.NSpace) {
				t.Fatalf("got result %+v, expected was %+v", got, tc.NSpace)
			}
			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestService_UpdateNamespace(t *testing.T) {
	type testCase struct {
		Description string
		NSpace      *namespace.Namespace
		Setup       func(*mocks.Repository, *mocks.Encryptor, *mocks.ConfigSyncer, testCase)
		Err         error
	}
	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "should return error if namespace is nil",
				Err:         errors.New("incoming namespace is empty"),
			},
			{
				Description: "should return error if provider service return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(nil, errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if decrypt return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if encrypt return error caused credential is not in json",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"invalid": make(chan int),
					},
				},
				Err: errors.New("json: unsupported type: chan int"),
			},
			{
				Description: "should return error if encrypt return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("", errors.New("some error"))
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if update repository error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(errors.New("some error"))
					rr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error not found if update repository return not found error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(namespace.NotFoundError{})
					rr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("namespace.NotFoundError")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("namespace not found"),
			},
			{
				Description: "should return error not found if update repository return relation error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(namespace.ErrRelation)
					rr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("provider id does not exist"),
			},
			{
				Description: "should return error conflict if update repository return error duplicate",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(namespace.ErrDuplicate)
					rr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("urn and provider pair already exist"),
			},
			{
				Description: "should return error if sync config return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(nil)
					cs.EXPECT().SyncRuntimeConfig(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("provider.Provider")).Return(nil, errors.New("some error"))
					rr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return error if update repository success, sync success, and update labels return error",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(nil)
					cs.EXPECT().SyncRuntimeConfig(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("provider.Provider")).Return(map[string]string{"k": "v"}, nil)
					rr.EXPECT().UpdateLabels(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(errors.New("some error"))
					rr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: errors.New("some error"),
			},
			{
				Description: "should return nil error if update repository success, sync, and update labels success",
				Setup: func(rr *mocks.Repository, e *mocks.Encryptor, cs *mocks.ConfigSyncer, tc testCase) {
					rr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&namespace.EncryptedNamespace{Namespace: &namespace.Namespace{Provider: provider.Provider{Type: testProviderType}}}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("secret.MaskableString")).Return("{ \"key\": \"value\" }", nil)
					e.EXPECT().Encrypt(mock.AnythingOfType("secret.MaskableString")).Return("some-ciphertext", nil)
					rr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(ctx)
					rr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), &namespace.EncryptedNamespace{
						Namespace:        tc.NSpace,
						CredentialString: "some-ciphertext",
					}).Return(nil)
					cs.EXPECT().SyncRuntimeConfig(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("provider.Provider")).Return(map[string]string{"k": "v"}, nil)
					rr.EXPECT().UpdateLabels(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64"), mock.AnythingOfType("map[string]string")).Return(nil)
					rr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]any{
						"credential": "value",
					},
				},
				Err: nil,
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock     = new(mocks.Repository)
				encryptorMock      = new(mocks.Encryptor)
				providerPluginMock = new(mocks.ConfigSyncer)
			)
			svc := namespace.NewService(encryptorMock, repositoryMock, nil,
				map[string]namespace.ConfigSyncer{
					testProviderType: providerPluginMock,
				},
			)

			if tc.Setup != nil {
				tc.Setup(repositoryMock, encryptorMock, providerPluginMock, tc)
			}

			err := svc.Update(ctx, tc.NSpace)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}

			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestDeleteNamespace(t *testing.T) {
	ctx := context.TODO()
	namespaceID := uint64(10)

	t.Run("should call repository Delete method and return nil if no error", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := namespace.NewService(nil, repositoryMock, nil, nil)
		repositoryMock.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), namespaceID).Return(nil).Once()
		err := dummyService.Delete(ctx, namespaceID)
		assert.Nil(t, err)
		repositoryMock.AssertExpectations(t)
	})

	t.Run("should call repository Delete method and return error if any", func(t *testing.T) {
		repositoryMock := &mocks.Repository{}
		dummyService := namespace.NewService(nil, repositoryMock, nil, nil)
		repositoryMock.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), namespaceID).Return(errors.New("random error")).Once()
		err := dummyService.Delete(ctx, namespaceID)
		assert.EqualError(t, err, "random error")
		repositoryMock.AssertExpectations(t)
	})
}
