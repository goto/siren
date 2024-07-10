"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[9783],{69122:(e,a,t)=>{t.r(a),t.d(a,{assets:()=>g,contentTitle:()=>d,default:()=>f,frontMatter:()=>c,metadata:()=>y,toc:()=>u});var s=t(58168),i=(t(96540),t(15680)),n=t(54754),r=t(34366),p=(t(13185),t(26612)),m=t(30774),l=t(47335),o=(t(38454),t(19365));const c={id:"siren-service-create-namespace",title:"create a namespace",description:"create a namespace",sidebar_label:"create a namespace",hide_title:!0,hide_table_of_contents:!0,api:{operationId:"SirenService_CreateNamespace",responses:{200:{description:"A successful response.",content:{"application/json":{schema:{type:"object",properties:{id:{type:"string",format:"uint64"}}}}}},default:{description:"An unexpected error response.",content:{"application/json":{schema:{type:"object",properties:{code:{type:"integer",format:"int32"},message:{type:"string"},details:{type:"array",items:{type:"object",properties:{"@type":{type:"string"}},additionalProperties:{}}}}}}}}},requestBody:{content:{"application/json":{schema:{type:"object",properties:{name:{type:"string"},urn:{type:"string"},provider:{type:"string",format:"uint64"},credentials:{type:"object"},labels:{type:"object",additionalProperties:{type:"string"}},created_at:{type:"string",format:"date-time"},updated_at:{type:"string",format:"date-time"}}}}},required:!0},tags:["Namespace"],description:"create a namespace",method:"post",path:"/v1beta1/namespaces",jsonRequestBodyExample:{name:"string",urn:"string",provider:"string",credentials:{},labels:{},created_at:"2024-07-09T13:51:06.041Z",updated_at:"2024-07-09T13:51:06.041Z"},info:{title:"Siren APIs",description:"Documentation of our Siren API with gRPC and\ngRPC-Gateway.",version:"0.6"},postman:{name:"create a namespace",description:{type:"text/plain"},url:{path:["v1beta1","namespaces"],host:["{{baseUrl}}"],query:[],variable:[]},header:[{key:"Content-Type",value:"application/json"},{key:"Accept",value:"application/json"}],method:"POST",body:{mode:"raw",raw:'""',options:{raw:{language:"json"}}}}},sidebar_class_name:"post api-method",info_path:"docs/apis/siren-apis",custom_edit_url:null},d=void 0,y={unversionedId:"apis/siren-service-create-namespace",id:"apis/siren-service-create-namespace",title:"create a namespace",description:"create a namespace",source:"@site/docs/apis/siren-service-create-namespace.api.mdx",sourceDirName:"apis",slug:"/apis/siren-service-create-namespace",permalink:"/siren/docs/apis/siren-service-create-namespace",draft:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"siren-service-create-namespace",title:"create a namespace",description:"create a namespace",sidebar_label:"create a namespace",hide_title:!0,hide_table_of_contents:!0,api:{operationId:"SirenService_CreateNamespace",responses:{200:{description:"A successful response.",content:{"application/json":{schema:{type:"object",properties:{id:{type:"string",format:"uint64"}}}}}},default:{description:"An unexpected error response.",content:{"application/json":{schema:{type:"object",properties:{code:{type:"integer",format:"int32"},message:{type:"string"},details:{type:"array",items:{type:"object",properties:{"@type":{type:"string"}},additionalProperties:{}}}}}}}}},requestBody:{content:{"application/json":{schema:{type:"object",properties:{name:{type:"string"},urn:{type:"string"},provider:{type:"string",format:"uint64"},credentials:{type:"object"},labels:{type:"object",additionalProperties:{type:"string"}},created_at:{type:"string",format:"date-time"},updated_at:{type:"string",format:"date-time"}}}}},required:!0},tags:["Namespace"],description:"create a namespace",method:"post",path:"/v1beta1/namespaces",jsonRequestBodyExample:{name:"string",urn:"string",provider:"string",credentials:{},labels:{},created_at:"2024-07-09T13:51:06.041Z",updated_at:"2024-07-09T13:51:06.041Z"},info:{title:"Siren APIs",description:"Documentation of our Siren API with gRPC and\ngRPC-Gateway.",version:"0.6"},postman:{name:"create a namespace",description:{type:"text/plain"},url:{path:["v1beta1","namespaces"],host:["{{baseUrl}}"],query:[],variable:[]},header:[{key:"Content-Type",value:"application/json"},{key:"Accept",value:"application/json"}],method:"POST",body:{mode:"raw",raw:'""',options:{raw:{language:"json"}}}}},sidebar_class_name:"post api-method",info_path:"docs/apis/siren-apis",custom_edit_url:null},sidebar:"docsSidebar",previous:{title:"list namespaces",permalink:"/siren/docs/apis/siren-service-list-namespaces"},next:{title:"get a namespace",permalink:"/siren/docs/apis/siren-service-get-namespace"}},g={},u=[{value:"create a namespace",id:"create-a-namespace",level:2}],h={toc:u},b="wrapper";function f(e){let{components:a,...t}=e;return(0,i.yg)(b,(0,s.A)({},h,t,{components:a,mdxType:"MDXLayout"}),(0,i.yg)("h2",{id:"create-a-namespace"},"create a namespace"),(0,i.yg)("p",null,"create a namespace"),(0,i.yg)(r.A,{mdxType:"MimeTabs"},(0,i.yg)(o.A,{label:"application/json",value:"application/json-schema",mdxType:"TabItem"},(0,i.yg)("details",{style:{},"data-collapsed":!1,open:!0},(0,i.yg)("summary",{style:{textAlign:"left"}},(0,i.yg)("strong",null,"Request Body"),(0,i.yg)("strong",{style:{fontSize:"var(--ifm-code-font-size)",color:"var(--openapi-required)"}}," required")),(0,i.yg)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.yg)("ul",{style:{marginLeft:"1rem"}},(0,i.yg)(m.A,{collapsible:!1,name:"name",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!1,name:"urn",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!1,name:"provider",required:!1,schemaName:"uint64",qualifierMessage:void 0,schema:{type:"string",format:"uint64"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!1,name:"credentials",required:!1,schemaName:"object",qualifierMessage:void 0,schema:{type:"object"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,i.yg)("details",{style:{}},(0,i.yg)("summary",{style:{}},(0,i.yg)("strong",null,"labels"),(0,i.yg)("span",{style:{opacity:"0.6"}}," object")),(0,i.yg)("div",{style:{marginLeft:"1rem"}},(0,i.yg)(m.A,{name:"property name*",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},collapsible:!1,discriminator:!1,mdxType:"SchemaItem"})))),(0,i.yg)(m.A,{collapsible:!1,name:"created_at",required:!1,schemaName:"date-time",qualifierMessage:void 0,schema:{type:"string",format:"date-time"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!1,name:"updated_at",required:!1,schemaName:"date-time",qualifierMessage:void 0,schema:{type:"string",format:"date-time"},mdxType:"SchemaItem"}))))),(0,i.yg)("div",null,(0,i.yg)(n.A,{mdxType:"ApiTabs"},(0,i.yg)(o.A,{label:"200",value:"200",mdxType:"TabItem"},(0,i.yg)("div",null,(0,i.yg)("p",null,"A successful response.")),(0,i.yg)("div",null,(0,i.yg)(r.A,{schemaType:"response",mdxType:"MimeTabs"},(0,i.yg)(o.A,{label:"application/json",value:"application/json",mdxType:"TabItem"},(0,i.yg)(l.A,{mdxType:"SchemaTabs"},(0,i.yg)(o.A,{label:"Schema",value:"Schema",mdxType:"TabItem"},(0,i.yg)("details",{style:{},"data-collapsed":!1,open:!0},(0,i.yg)("summary",{style:{textAlign:"left"}},(0,i.yg)("strong",null,"Schema")),(0,i.yg)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.yg)("ul",{style:{marginLeft:"1rem"}},(0,i.yg)(m.A,{collapsible:!1,name:"id",required:!1,schemaName:"uint64",qualifierMessage:void 0,schema:{type:"string",format:"uint64"},mdxType:"SchemaItem"})))),(0,i.yg)(o.A,{label:"Example (from schema)",value:"Example (from schema)",mdxType:"TabItem"},(0,i.yg)(p.A,{responseExample:'{\n  "id": "string"\n}',language:"json",mdxType:"ResponseSamples"}))))))),(0,i.yg)(o.A,{label:"default",value:"default",mdxType:"TabItem"},(0,i.yg)("div",null,(0,i.yg)("p",null,"An unexpected error response.")),(0,i.yg)("div",null,(0,i.yg)(r.A,{schemaType:"response",mdxType:"MimeTabs"},(0,i.yg)(o.A,{label:"application/json",value:"application/json",mdxType:"TabItem"},(0,i.yg)(l.A,{mdxType:"SchemaTabs"},(0,i.yg)(o.A,{label:"Schema",value:"Schema",mdxType:"TabItem"},(0,i.yg)("details",{style:{},"data-collapsed":!1,open:!0},(0,i.yg)("summary",{style:{textAlign:"left"}},(0,i.yg)("strong",null,"Schema")),(0,i.yg)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.yg)("ul",{style:{marginLeft:"1rem"}},(0,i.yg)(m.A,{collapsible:!1,name:"code",required:!1,schemaName:"int32",qualifierMessage:void 0,schema:{type:"integer",format:"int32"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!1,name:"message",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,i.yg)(m.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,i.yg)("details",{style:{}},(0,i.yg)("summary",{style:{}},(0,i.yg)("strong",null,"details"),(0,i.yg)("span",{style:{opacity:"0.6"}}," object[]")),(0,i.yg)("div",{style:{marginLeft:"1rem"}},(0,i.yg)("li",null,(0,i.yg)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"}},"Array [")),(0,i.yg)(m.A,{collapsible:!1,name:"@type",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,i.yg)("li",null,(0,i.yg)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"}},"]")))))))),(0,i.yg)(o.A,{label:"Example (from schema)",value:"Example (from schema)",mdxType:"TabItem"},(0,i.yg)(p.A,{responseExample:'{\n  "code": 0,\n  "message": "string",\n  "details": [\n    {\n      "@type": "string"\n    }\n  ]\n}',language:"json",mdxType:"ResponseSamples"}))))))))))}f.isMDXComponent=!0}}]);