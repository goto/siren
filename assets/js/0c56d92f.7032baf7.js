"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[7704],{74557:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>u,frontMatter:()=>a,metadata:()=>s,toc:()=>c});var i=r(58168),n=(r(96540),r(15680));const a={},o="CortexMetrics",s={unversionedId:"providers/cortexmetrics",id:"providers/cortexmetrics",title:"CortexMetrics",description:"|||",source:"@site/docs/providers/cortexmetrics.md",sourceDirName:"providers",slug:"/providers/cortexmetrics",permalink:"/siren/docs/providers/cortexmetrics",draft:!1,editUrl:"https://github.com/goto/siren/edit/master/docs/docs/providers/cortexmetrics.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Job",permalink:"/siren/docs/guides/job"},next:{title:"Slack",permalink:"/siren/docs/receivers/slack"}},l={},c=[{value:"Multi-tenancy",id:"multi-tenancy",level:2},{value:"Server Configuration",id:"server-configuration",level:2}],d={toc:c},p="wrapper";function u(e){let{components:t,...r}=e;return(0,n.yg)(p,(0,i.A)({},d,r,{components:t,mdxType:"MDXLayout"}),(0,n.yg)("h1",{id:"cortexmetrics"},"CortexMetrics"),(0,n.yg)("table",null,(0,n.yg)("thead",{parentName:"table"},(0,n.yg)("tr",{parentName:"thead"},(0,n.yg)("th",{parentName:"tr",align:null}),(0,n.yg)("th",{parentName:"tr",align:null}))),(0,n.yg)("tbody",{parentName:"table"},(0,n.yg)("tr",{parentName:"tbody"},(0,n.yg)("td",{parentName:"tr",align:null},(0,n.yg)("strong",{parentName:"td"},"type")),(0,n.yg)("td",{parentName:"tr",align:null},(0,n.yg)("inlineCode",{parentName:"td"},"cortex"))))),(0,n.yg)("p",null,(0,n.yg)("a",{parentName:"p",href:"https://cortexmetrics.io/"},"CortexMetrics")," is a Horizontally scalable, highly available, multi-tenant, long term storage for Prometheus. It could run in multi-tenant scenario where it isolates data and queries from multiple different independent Prometheus sources in a single cluster, allowing untrusted parties to share the same cluster."),(0,n.yg)("p",null,"Similar with how prometheus works, CortexMetrics consumes metrics sent from other services, store, and evaluate the metrics with the configured rules. Alerts will be triggered and processed by CortexMetrics' alert manager and notifications will be sent to the designated sinks (webhook, slack, pagerduty, etc..). "),(0,n.yg)("p",null,"Since Siren handles all subscriptions and notifications routing, Siren configures CortexMetrics to send all alerts only to Siren webhook API."),(0,n.yg)("h2",{id:"multi-tenancy"},"Multi-tenancy"),(0,n.yg)("p",null,"Tenants in CortexMetrics are mapped to ",(0,n.yg)("a",{parentName:"p",href:"/siren/docs/guides/provider_and_namespace#namespace"},"Namespaces")," in Siren. To integrate multiple tenants, you need to create multiple namespaces for each tenant. Each tenant will have different configuration."),(0,n.yg)("h2",{id:"server-configuration"},"Server Configuration"),(0,n.yg)("p",null,"There is a generic CortexMetrics configuration in Siren server configuration that could be used to tune the CortexMetrics. The configuration will always be synchronized everytime a namespace in Siren is created or updated."),(0,n.yg)("p",null,"Here is a config that is part of the server configuration. Please note that the config will always be applied to all CortexMetrics registered in Siren and only synchronized when a namespace in Siren is created or updated. Siren server restart is also required to get the latest value update of these configs."),(0,n.yg)("pre",null,(0,n.yg)("code",{parentName:"pre",className:"language-yaml"},"...\nproviders:\n  cortex:\n    group_wait: 30s\n    webhook_base_api: http://localhost:8080/v1beta1/alerts/cortex\n...\n")),(0,n.yg)("ul",null,(0,n.yg)("li",{parentName:"ul"},"The ",(0,n.yg)("inlineCode",{parentName:"li"},"group_wait")," config usage is similar with the one in CortexMetrics alert manager ",(0,n.yg)("a",{parentName:"li",href:"https://prometheus.io/docs/alerting/latest/configuration/#example"},"configuration"),"."),(0,n.yg)("li",{parentName:"ul"},"The ",(0,n.yg)("inlineCode",{parentName:"li"},"webhook_base_api")," defined the base API that will be appended with ",(0,n.yg)("inlineCode",{parentName:"li"},"provider_id")," for each specific provider. If a namespace of provider with id ",(0,n.yg)("inlineCode",{parentName:"li"},"3")," is updated, Siren will configure the webhook receiver in CortexMetrics with this URL: ",(0,n.yg)("inlineCode",{parentName:"li"},"http://localhost:8080/v1beta1/alerts/cortex/3"),".")))}u.isMDXComponent=!0}}]);