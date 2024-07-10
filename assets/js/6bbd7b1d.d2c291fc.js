"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[8089],{93162:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>u,contentTitle:()=>d,default:()=>v,frontMatter:()=>y,metadata:()=>g,toc:()=>b});var a=i(58168),s=(i(96540),i(15680)),r=i(54754),n=i(34366),p=i(13185),m=i(26612),o=i(30774),l=i(47335),c=(i(38454),i(19365));const y={id:"siren-service-get-subscription",title:"Get a subscription",description:"Get a subscription",sidebar_label:"Get a subscription",hide_title:!0,hide_table_of_contents:!0,api:{operationId:"SirenService_GetSubscription",responses:{200:{description:"A successful response.",content:{"application/json":{schema:{type:"object",properties:{subscription:{type:"object",properties:{id:{type:"string",format:"uint64"},urn:{type:"string"},namespace:{type:"string",format:"uint64"},receivers:{type:"array",items:{type:"object",properties:{id:{type:"string",format:"uint64"},configuration:{type:"object"},subscription_receiver_labels:{type:"object",additionalProperties:{type:"string"}}}}},match:{type:"object",additionalProperties:{type:"string"}},created_at:{type:"string",format:"date-time"},updated_at:{type:"string",format:"date-time"},metadata:{type:"object"},created_by:{type:"string"},updated_by:{type:"string"}}}}}}}},default:{description:"An unexpected error response.",content:{"application/json":{schema:{type:"object",properties:{code:{type:"integer",format:"int32"},message:{type:"string"},details:{type:"array",items:{type:"object",properties:{"@type":{type:"string"}},additionalProperties:{}}}}}}}}},parameters:[{name:"id",in:"path",required:!0,schema:{type:"string",format:"uint64"}}],tags:["Subscription"],description:"Get a subscription",method:"get",path:"/v1beta1/subscriptions/{id}",info:{title:"Siren APIs",description:"Documentation of our Siren API with gRPC and\ngRPC-Gateway.",version:"0.6"},postman:{name:"Get a subscription",description:{type:"text/plain"},url:{path:["v1beta1","subscriptions",":id"],host:["{{baseUrl}}"],query:[],variable:[{disabled:!1,description:{content:"(Required) ",type:"text/plain"},type:"any",value:"",key:"id"}]},header:[{key:"Accept",value:"application/json"}],method:"GET"}},sidebar_class_name:"get api-method",info_path:"docs/apis/siren-apis",custom_edit_url:null},d=void 0,g={unversionedId:"apis/siren-service-get-subscription",id:"apis/siren-service-get-subscription",title:"Get a subscription",description:"Get a subscription",source:"@site/docs/apis/siren-service-get-subscription.api.mdx",sourceDirName:"apis",slug:"/apis/siren-service-get-subscription",permalink:"/siren/docs/apis/siren-service-get-subscription",draft:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"siren-service-get-subscription",title:"Get a subscription",description:"Get a subscription",sidebar_label:"Get a subscription",hide_title:!0,hide_table_of_contents:!0,api:{operationId:"SirenService_GetSubscription",responses:{200:{description:"A successful response.",content:{"application/json":{schema:{type:"object",properties:{subscription:{type:"object",properties:{id:{type:"string",format:"uint64"},urn:{type:"string"},namespace:{type:"string",format:"uint64"},receivers:{type:"array",items:{type:"object",properties:{id:{type:"string",format:"uint64"},configuration:{type:"object"},subscription_receiver_labels:{type:"object",additionalProperties:{type:"string"}}}}},match:{type:"object",additionalProperties:{type:"string"}},created_at:{type:"string",format:"date-time"},updated_at:{type:"string",format:"date-time"},metadata:{type:"object"},created_by:{type:"string"},updated_by:{type:"string"}}}}}}}},default:{description:"An unexpected error response.",content:{"application/json":{schema:{type:"object",properties:{code:{type:"integer",format:"int32"},message:{type:"string"},details:{type:"array",items:{type:"object",properties:{"@type":{type:"string"}},additionalProperties:{}}}}}}}}},parameters:[{name:"id",in:"path",required:!0,schema:{type:"string",format:"uint64"}}],tags:["Subscription"],description:"Get a subscription",method:"get",path:"/v1beta1/subscriptions/{id}",info:{title:"Siren APIs",description:"Documentation of our Siren API with gRPC and\ngRPC-Gateway.",version:"0.6"},postman:{name:"Get a subscription",description:{type:"text/plain"},url:{path:["v1beta1","subscriptions",":id"],host:["{{baseUrl}}"],query:[],variable:[{disabled:!1,description:{content:"(Required) ",type:"text/plain"},type:"any",value:"",key:"id"}]},header:[{key:"Accept",value:"application/json"}],method:"GET"}},sidebar_class_name:"get api-method",info_path:"docs/apis/siren-apis",custom_edit_url:null},sidebar:"docsSidebar",previous:{title:"Create a subscription",permalink:"/siren/docs/apis/siren-service-create-subscription"},next:{title:"Delete a subscription",permalink:"/siren/docs/apis/siren-service-delete-subscription"}},u={},b=[{value:"Get a subscription",id:"get-a-subscription",level:2}],h={toc:b},f="wrapper";function v(e){let{components:t,...i}=e;return(0,s.yg)(f,(0,a.A)({},h,i,{components:t,mdxType:"MDXLayout"}),(0,s.yg)("h2",{id:"get-a-subscription"},"Get a subscription"),(0,s.yg)("p",null,"Get a subscription"),(0,s.yg)("details",{style:{marginBottom:"1rem"},"data-collapsed":!1,open:!0},(0,s.yg)("summary",{style:{}},(0,s.yg)("strong",null,"Path Parameters")),(0,s.yg)("div",null,(0,s.yg)("ul",null,(0,s.yg)(p.A,{className:"paramsItem",param:{name:"id",in:"path",required:!0,schema:{type:"string",format:"uint64"}},mdxType:"ParamsItem"})))),(0,s.yg)("div",null,(0,s.yg)(r.A,{mdxType:"ApiTabs"},(0,s.yg)(c.A,{label:"200",value:"200",mdxType:"TabItem"},(0,s.yg)("div",null,(0,s.yg)("p",null,"A successful response.")),(0,s.yg)("div",null,(0,s.yg)(n.A,{schemaType:"response",mdxType:"MimeTabs"},(0,s.yg)(c.A,{label:"application/json",value:"application/json",mdxType:"TabItem"},(0,s.yg)(l.A,{mdxType:"SchemaTabs"},(0,s.yg)(c.A,{label:"Schema",value:"Schema",mdxType:"TabItem"},(0,s.yg)("details",{style:{},"data-collapsed":!1,open:!0},(0,s.yg)("summary",{style:{textAlign:"left"}},(0,s.yg)("strong",null,"Schema")),(0,s.yg)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,s.yg)("ul",{style:{marginLeft:"1rem"}},(0,s.yg)(o.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,s.yg)("details",{style:{}},(0,s.yg)("summary",{style:{}},(0,s.yg)("strong",null,"subscription"),(0,s.yg)("span",{style:{opacity:"0.6"}}," object")),(0,s.yg)("div",{style:{marginLeft:"1rem"}},(0,s.yg)(o.A,{collapsible:!1,name:"id",required:!1,schemaName:"uint64",qualifierMessage:void 0,schema:{type:"string",format:"uint64"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"urn",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"namespace",required:!1,schemaName:"uint64",qualifierMessage:void 0,schema:{type:"string",format:"uint64"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,s.yg)("details",{style:{}},(0,s.yg)("summary",{style:{}},(0,s.yg)("strong",null,"receivers"),(0,s.yg)("span",{style:{opacity:"0.6"}}," object[]")),(0,s.yg)("div",{style:{marginLeft:"1rem"}},(0,s.yg)("li",null,(0,s.yg)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"}},"Array [")),(0,s.yg)(o.A,{collapsible:!1,name:"id",required:!1,schemaName:"uint64",qualifierMessage:void 0,schema:{type:"string",format:"uint64"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"configuration",required:!1,schemaName:"object",qualifierMessage:void 0,schema:{type:"object"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,s.yg)("details",{style:{}},(0,s.yg)("summary",{style:{}},(0,s.yg)("strong",null,"subscription_receiver_labels"),(0,s.yg)("span",{style:{opacity:"0.6"}}," object")),(0,s.yg)("div",{style:{marginLeft:"1rem"}},(0,s.yg)(o.A,{name:"property name*",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},collapsible:!1,discriminator:!1,mdxType:"SchemaItem"})))),(0,s.yg)("li",null,(0,s.yg)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"}},"]"))))),(0,s.yg)(o.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,s.yg)("details",{style:{}},(0,s.yg)("summary",{style:{}},(0,s.yg)("strong",null,"match"),(0,s.yg)("span",{style:{opacity:"0.6"}}," object")),(0,s.yg)("div",{style:{marginLeft:"1rem"}},(0,s.yg)(o.A,{name:"property name*",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},collapsible:!1,discriminator:!1,mdxType:"SchemaItem"})))),(0,s.yg)(o.A,{collapsible:!1,name:"created_at",required:!1,schemaName:"date-time",qualifierMessage:void 0,schema:{type:"string",format:"date-time"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"updated_at",required:!1,schemaName:"date-time",qualifierMessage:void 0,schema:{type:"string",format:"date-time"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"metadata",required:!1,schemaName:"object",qualifierMessage:void 0,schema:{type:"object"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"created_by",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"updated_by",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}))))))),(0,s.yg)(c.A,{label:"Example (from schema)",value:"Example (from schema)",mdxType:"TabItem"},(0,s.yg)(m.A,{responseExample:'{\n  "subscription": {\n    "id": "string",\n    "urn": "string",\n    "namespace": "string",\n    "receivers": [\n      {\n        "id": "string",\n        "configuration": {},\n        "subscription_receiver_labels": {}\n      }\n    ],\n    "match": {},\n    "created_at": "2024-07-09T13:51:06.083Z",\n    "updated_at": "2024-07-09T13:51:06.083Z",\n    "metadata": {},\n    "created_by": "string",\n    "updated_by": "string"\n  }\n}',language:"json",mdxType:"ResponseSamples"}))))))),(0,s.yg)(c.A,{label:"default",value:"default",mdxType:"TabItem"},(0,s.yg)("div",null,(0,s.yg)("p",null,"An unexpected error response.")),(0,s.yg)("div",null,(0,s.yg)(n.A,{schemaType:"response",mdxType:"MimeTabs"},(0,s.yg)(c.A,{label:"application/json",value:"application/json",mdxType:"TabItem"},(0,s.yg)(l.A,{mdxType:"SchemaTabs"},(0,s.yg)(c.A,{label:"Schema",value:"Schema",mdxType:"TabItem"},(0,s.yg)("details",{style:{},"data-collapsed":!1,open:!0},(0,s.yg)("summary",{style:{textAlign:"left"}},(0,s.yg)("strong",null,"Schema")),(0,s.yg)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,s.yg)("ul",{style:{marginLeft:"1rem"}},(0,s.yg)(o.A,{collapsible:!1,name:"code",required:!1,schemaName:"int32",qualifierMessage:void 0,schema:{type:"integer",format:"int32"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!1,name:"message",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,s.yg)(o.A,{collapsible:!0,className:"schemaItem",mdxType:"SchemaItem"},(0,s.yg)("details",{style:{}},(0,s.yg)("summary",{style:{}},(0,s.yg)("strong",null,"details"),(0,s.yg)("span",{style:{opacity:"0.6"}}," object[]")),(0,s.yg)("div",{style:{marginLeft:"1rem"}},(0,s.yg)("li",null,(0,s.yg)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"}},"Array [")),(0,s.yg)(o.A,{collapsible:!1,name:"@type",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"},mdxType:"SchemaItem"}),(0,s.yg)("li",null,(0,s.yg)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"}},"]")))))))),(0,s.yg)(c.A,{label:"Example (from schema)",value:"Example (from schema)",mdxType:"TabItem"},(0,s.yg)(m.A,{responseExample:'{\n  "code": 0,\n  "message": "string",\n  "details": [\n    {\n      "@type": "string"\n    }\n  ]\n}',language:"json",mdxType:"ResponseSamples"}))))))))))}v.isMDXComponent=!0}}]);