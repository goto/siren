package postgres_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/odpf/siren/mocks"
	"github.com/odpf/siren/store/model"
	"github.com/odpf/siren/store/postgres"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type AlertsRepositoryTestSuite struct {
	suite.Suite
	sqldb      *sql.DB
	dbmock     sqlmock.Sqlmock
	repository model.AlertRepository
}

func (s *AlertsRepositoryTestSuite) SetupTest() {
	db, mock, _ := mocks.NewStore()
	s.sqldb, _ = db.DB()
	s.dbmock = mock
	s.repository = postgres.NewRepository(db)
}

func (s *AlertsRepositoryTestSuite) TearDownTest() {
	s.sqldb.Close()
}

func (s *AlertsRepositoryTestSuite) TestGet() {
	timenow := time.Now()
	s.Run("should fetch matching alert history objects", func() {
		expectedQuery := regexp.QuoteMeta(`select * from alerts where resource_name = 'foo' AND provider_id = '1' AND triggered_at BETWEEN to_timestamp('0') AND to_timestamp('1000')`)
		expectedAlert := model.Alert{
			Id: 1, ProviderId: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "20",
			Rule: "bar", TriggeredAt: timenow, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}
		expectedAlerts := []model.Alert{expectedAlert}
		expectedRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "resource_name", "provider_id",
			"severity", "metric_name", "metric_value", "rule", "triggered_at"}).
			AddRow(expectedAlert.Id, expectedAlert.CreatedAt,
				expectedAlert.UpdatedAt, expectedAlert.ResourceName, expectedAlert.ProviderId, expectedAlert.Severity,
				expectedAlert.MetricName, expectedAlert.MetricValue, expectedAlert.Rule, expectedAlert.TriggeredAt)
		s.dbmock.ExpectQuery(expectedQuery).WillReturnRows(expectedRows)

		actualAlerts, err := s.repository.Get("foo", 1, 0, 1000)
		s.Equal(expectedAlerts, actualAlerts)
		s.Nil(err)
	})

	s.Run("should return error if any in fetching alert history objects", func() {
		expectedQuery := regexp.QuoteMeta(`select * from alerts where resource_name = 'foo' AND provider_id = '1' AND triggered_at BETWEEN to_timestamp('0') AND to_timestamp('1000')`)
		s.dbmock.ExpectQuery(expectedQuery).WillReturnError(errors.New("random error"))
		actualAlerts, err := s.repository.Get("foo", 1, 0, 1000)
		s.Nil(actualAlerts)
		s.EqualError(err, "random error")
	})
}

func (s *AlertsRepositoryTestSuite) TestCreate() {
	timenow := time.Now()
	s.Run("should create alert object", func() {
		insertQuery := regexp.QuoteMeta(`INSERT INTO "alerts" ("provider_id","resource_name","metric_name","metric_value","severity","rule","triggered_at","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`)
		expectedAlerts := &model.Alert{
			Id: 1, ProviderId: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "20",
			Rule: "bar", TriggeredAt: timenow, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}
		s.dbmock.ExpectQuery(insertQuery).WithArgs(expectedAlerts.ProviderId,
			expectedAlerts.ResourceName, expectedAlerts.MetricName, expectedAlerts.MetricValue, expectedAlerts.Severity,
			expectedAlerts.Rule, expectedAlerts.TriggeredAt, expectedAlerts.CreatedAt, expectedAlerts.UpdatedAt,
			expectedAlerts.Id).
			WillReturnRows(sqlmock.NewRows(nil))
		actualAlert, err := s.repository.Create(expectedAlerts)
		s.Equal(expectedAlerts, actualAlert)
		s.Nil(err)
	})

	s.Run("should return error in alert history creation", func() {
		insertQuery := regexp.QuoteMeta(`INSERT INTO "alerts" ("provider_id","resource_name","metric_name","metric_value","severity","rule","triggered_at","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`)
		expectedAlerts := &model.Alert{
			Id: 1, ProviderId: 1, ResourceName: "foo", Severity: "CRITICAL", MetricName: "baz", MetricValue: "20",
			Rule: "bar", TriggeredAt: timenow, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}
		s.dbmock.ExpectQuery(insertQuery).WithArgs(expectedAlerts.ProviderId,
			expectedAlerts.ResourceName, expectedAlerts.MetricName, expectedAlerts.MetricValue, expectedAlerts.Severity,
			expectedAlerts.Rule, expectedAlerts.TriggeredAt, expectedAlerts.CreatedAt, expectedAlerts.UpdatedAt,
			expectedAlerts.Id).
			WillReturnError(errors.New("random error"))
		actualAlert, err := s.repository.Create(expectedAlerts)
		s.Nil(actualAlert)
		s.EqualError(err, "random error")
	})
}

func TestAlertsRepository(t *testing.T) {
	suite.Run(t, new(AlertsRepositoryTestSuite))
}