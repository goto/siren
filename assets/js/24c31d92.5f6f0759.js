(self.webpackChunksiren=self.webpackChunksiren||[]).push([[439],{1410:(e,t,a)=>{const n=a(7694),l=a(3618);e.exports={title:"Siren",tagline:"Universal data observability toolkit",url:"https://odpf.github.io",baseUrl:"/siren/",onBrokenLinks:"throw",onBrokenMarkdownLinks:"warn",favicon:"img/favicon.ico",organizationName:"odpf",projectName:"siren",customFields:{apiVersion:"v1beta1",defaultHost:"http://localhost:8080"},presets:[["@docusaurus/preset-classic",{docs:{sidebarPath:6679,editUrl:"https://github.com/odpf/siren/edit/master/docs/",sidebarCollapsed:!0,breadcrumbs:!1},blog:!1,theme:{customCss:[5308,2295]},gtag:{trackingID:"G-EPXDLH6V72"}}]],themeConfig:{colorMode:{defaultMode:"light",respectPrefersColorScheme:!0},navbar:{title:"Siren",logo:{src:"img/logo.svg"},hideOnScroll:!0,items:[{type:"doc",docId:"introduction",position:"right",label:"Docs"},{to:"docs/support",label:"Support",position:"right"},{href:"https://bit.ly/2RzPbtn",position:"right",className:"header-slack-link"},{href:"https://github.com/odpf/siren",className:"navbar-item-github",position:"right"}]},footer:{style:"light",links:[]},prism:{theme:n,darkTheme:l},announcementBar:{id:"star-repo",content:'\u2b50\ufe0f If you like Siren, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/odpf/siren">GitHub</a>! \u2b50',backgroundColor:"#222",textColor:"#eee",isCloseable:!0}}}},5162:(e,t,a)=>{"use strict";a.d(t,{Z:()=>i});var n=a(7294),l=a(4334);const r="tabItem_Ymn6";function i(e){let{children:t,hidden:a,className:i}=e;return n.createElement("div",{role:"tabpanel",className:(0,l.Z)(r,i),hidden:a},t)}},5488:(e,t,a)=>{"use strict";a.d(t,{Z:()=>m});var n=a(3117),l=a(7294),r=a(4334),i=a(2389),o=a(7392),s=a(7094),p=a(2466);const u="tabList__CuJ",d="tabItem_LNqP";function c(e){var t;const{lazy:a,block:i,defaultValue:c,values:m,groupId:g,className:b}=e,h=l.Children.map(e.children,(e=>{if((0,l.isValidElement)(e)&&"value"in e.props)return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})),y=m??h.map((e=>{let{props:{value:t,label:a,attributes:n}}=e;return{value:t,label:a,attributes:n}})),f=(0,o.l)(y,((e,t)=>e.value===t.value));if(f.length>0)throw new Error(`Docusaurus error: Duplicate values "${f.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`);const k=null===c?c:c??(null==(t=h.find((e=>e.props.default)))?void 0:t.props.value)??h[0].props.value;if(null!==k&&!y.some((e=>e.value===k)))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${k}" but none of its children has the corresponding value. Available values are: ${y.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);const{tabGroupChoices:v,setTabGroupChoices:N}=(0,s.U)(),[T,C]=(0,l.useState)(k),I=[],{blockElementScrollPositionUntilNextRender:w}=(0,p.o5)();if(null!=g){const e=v[g];null!=e&&e!==T&&y.some((t=>t.value===e))&&C(e)}const x=e=>{const t=e.currentTarget,a=I.indexOf(t),n=y[a].value;n!==T&&(w(t),C(n),null!=g&&N(g,String(n)))},P=e=>{var t;let a=null;switch(e.key){case"Enter":x(e);break;case"ArrowRight":{const t=I.indexOf(e.currentTarget)+1;a=I[t]??I[0];break}case"ArrowLeft":{const t=I.indexOf(e.currentTarget)-1;a=I[t]??I[I.length-1];break}}null==(t=a)||t.focus()};return l.createElement("div",{className:(0,r.Z)("tabs-container",u)},l.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,r.Z)("tabs",{"tabs--block":i},b)},y.map((e=>{let{value:t,label:a,attributes:i}=e;return l.createElement("li",(0,n.Z)({role:"tab",tabIndex:T===t?0:-1,"aria-selected":T===t,key:t,ref:e=>I.push(e),onKeyDown:P,onClick:x},i,{className:(0,r.Z)("tabs__item",d,null==i?void 0:i.className,{"tabs__item--active":T===t})}),a??t)}))),a?(0,l.cloneElement)(h.filter((e=>e.props.value===T))[0],{className:"margin-top--md"}):l.createElement("div",{className:"margin-top--md"},h.map(((e,t)=>(0,l.cloneElement)(e,{key:t,hidden:e.props.value!==T})))))}function m(e){const t=(0,i.Z)();return l.createElement(c,(0,n.Z)({key:String(t)},e))}},6679:e=>{e.exports={docsSidebar:["introduction","installation",{type:"category",label:"Tour",items:["tour/overview","tour/start_server","tour/registering_provider","tour/registering_receivers","tour/sending_notifications_to_receiver","tour/configuring_provider_alerting_rules","tour/subscribing_notifications"]},{type:"category",label:"Concepts",items:["concepts/overview","concepts/plugin","concepts/schema"]},{type:"category",label:"Guides",items:["guides/overview","guides/provider_and_namespace","guides/receiver","guides/subscription","guides/rule","guides/template","guides/alert_history","guides/notification","guides/deployment"]},{type:"category",label:"Contribute",items:["contribute/contribution","contribute/receiver","contribute/provider","contribute/release"]},{type:"category",label:"Reference",items:["reference/api","reference/server_configuration","reference/client_configuration","reference/receiver","reference/cli"]}]}},2160:(e,t,a)=>{"use strict";a.r(t),a.d(t,{apiVersion:()=>b,assets:()=>m,contentTitle:()=>d,default:()=>f,defaultHost:()=>h,frontMatter:()=>u,metadata:()=>c,toc:()=>g});var n=a(3117),l=(a(7294),a(3905)),r=a(5488),i=a(5162),o=a(6066),s=a(1410),p=a.n(s);const u={},d="Template",c={unversionedId:"guides/template",id:"guides/template",title:"Template",description:"Templates concept in Siren is used for abstraction. The usage is versatile enough to be used to abstract out rules and notification format. It utilises go-templates to provide data-driven templates for generating textual output. The template delimiter used is [[ and ]].",source:"@site/docs/guides/template.md",sourceDirName:"guides",slug:"/guides/template",permalink:"/siren/docs/guides/template",draft:!1,editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/template.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Rule",permalink:"/siren/docs/guides/rule"},next:{title:"Alert History",permalink:"/siren/docs/guides/alert_history"}},m={},g=[{value:"API interface",id:"api-interface",level:2},{value:"Template creation/update",id:"template-creationupdate",level:3},{value:"Terminology of the request body",id:"terminology-of-the-request-body",level:3},{value:"Fetching a template",id:"fetching-a-template",level:3},{value:"Deleting a template",id:"deleting-a-template",level:3},{value:"CLI interface",id:"cli-interface",level:2},{value:"Terminology",id:"terminology",level:3}],b=p().customFields.apiVersion,h=p().customFields.defaultHost,y={toc:g,apiVersion:b};function f(e){let{components:t,...a}=e;return(0,l.kt)("wrapper",(0,n.Z)({},y,a,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"template"},"Template"),(0,l.kt)("p",null,"Templates concept in Siren is used for abstraction. The usage is versatile enough to be used to abstract out rules and notification format. It utilises ",(0,l.kt)("a",{parentName:"p",href:"https://golang.org/pkg/text/template/"},"go-templates")," to provide data-driven templates for generating textual output. The template delimiter used is ",(0,l.kt)("inlineCode",{parentName:"p"},"[[")," and ",(0,l.kt)("inlineCode",{parentName:"p"},"]]"),"."),(0,l.kt)("p",null,"One can create templates using either HTTP APIs or CLI."),(0,l.kt)("h2",{id:"api-interface"},"API interface"),(0,l.kt)("h3",{id:"template-creationupdate"},"Template creation/update"),(0,l.kt)("p",null,"Templates can be created using Siren APIs. The below snippet describes an example."),(0,l.kt)(r.Z,{groupId:"api",mdxType:"Tabs"},(0,l.kt)(i.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren template upsert --file template.yaml\n"))),(0,l.kt)(i.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,l.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request PUT\n  --url ",h,"/",b,'/templates\n  --header \'content-type: application/json\'\n  --data-raw \'{\n    "name": "CPU",\n    "body": "- alert: CPUHighWarning\n  expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.warning]]\n  for: \'[[.for]]\'\n  labels:\n    severity: WARNING\n    team: \'[[ .team ]]\'\n  annotations:\n    dashboard: https://example.com\n    description: CPU has been above [[.warning]] for last [[.for]] {{ $labels.host }}\n- alert: CPUHighCritical\n  expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.critical]]\n  for: \'[[.for]]\'\n  labels:\n    severity: CRITICAL\n    team: \'[[ .team ]]\'\n  annotations:\n    dashboard: example.com\n    description: CPU has been above [[.critical]] for last [[.for]] {{ $labels.host }}\n",\n    "tags": [\n        "firehose",\n        "dagger"\n    ],\n    "variables": [\n        {\n            "name": "team",\n            "type": "string",\n            "default": "odpf",\n            "description": "Name of the team that owns the deployment"\n        },\n        {\n            "name": "for",\n            "type": "string",\n            "default": "10m",\n            "description": "For eg 5m, 2h; Golang duration format"\n        },\n        {\n            "name": "warning",\n            "type": "int",\n            "default": "85",\n            "description": ""\n        },\n        {\n            "name": "critical",\n            "type": "int",\n            "default": "95",\n            "description": ""\n        }\n    ]\n}\''))),(0,l.kt)("h3",{id:"terminology-of-the-request-body"},"Terminology of the request body"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Term"),(0,l.kt)("th",{parentName:"tr",align:null},"Description"),(0,l.kt)("th",{parentName:"tr",align:null},"Example/Default"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Name"),(0,l.kt)("td",{parentName:"tr",align:null},"Name of the template"),(0,l.kt)("td",{parentName:"tr",align:null},"CPUHigh")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Body"),(0,l.kt)("td",{parentName:"tr",align:null},"Array of rule body. The body can be templatized in go template format."),(0,l.kt)("td",{parentName:"tr",align:null},"See example above")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Variables"),(0,l.kt)("td",{parentName:"tr",align:null},"Array of variables that were templatized in the body with their data type, default value and description."),(0,l.kt)("td",{parentName:"tr",align:null},"See example above")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Tags"),(0,l.kt)("td",{parentName:"tr",align:null},"Array of resources/applications that can utilize this template"),(0,l.kt)("td",{parentName:"tr",align:null},"VM")))),(0,l.kt)("p",null,"The response body will look like this:"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "id": 38,\n  "CreatedAt": "2021-04-29T16:20:48.061862+05:30",\n  "UpdatedAt": "2021-04-29T16:22:19.978837+05:30",\n  "name": "CPU",\n  "body": "- alert: CPUHighWarning\\n  expr: avg by (host) (cpu_usage_user{cpu=\\"cpu-total\\"}) > [[.warning]]\\n  for: \'[[.for]]\'\\n  labels:\\n    severity: WARNING\\n    team: \'[[ .team ]]\'\\n  annotations:\\n    dashboard: https://example.com\\n    description: CPU has been above [[.warning]] for last [[.for]] {{ $labels.host }}\\n- alert: CPUHighCritical\\n  expr: avg by (host) (cpu_usage_user{cpu=\\"cpu-total\\"}) > [[.critical]]\\n  for: \'[[.for]]\'\\n  labels:\\n    severity: CRITICAL\\n    team: \'[[ .team ]]\'\\n  annotations:\\n    dashboard: example.com\\n    description: CPU has been above [[.critical]] for last [[.for]] {{ $labels.host }}\\n",\n  "tags": ["firehose", "dagger"],\n  "variables": [\n    {\n      "name": "team",\n      "type": "string",\n      "default": "odpf",\n      "description": "Name of the team that owns the deployment"\n    },\n    {\n      "name": "for",\n      "type": "string",\n      "default": "10m",\n      "description": "For eg 5m, 2h; Golang duration format"\n    },\n    {\n      "name": "warning",\n      "type": "int",\n      "default": "85",\n      "description": ""\n    },\n    {\n      "name": "critical",\n      "type": "int",\n      "default": "95",\n      "description": ""\n    }\n  ]\n}\n')),(0,l.kt)("h3",{id:"fetching-a-template"},"Fetching a template"),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Fetching by Name")),(0,l.kt)("p",null,"Here is an example to fetch a template using name."),(0,l.kt)(r.Z,{groupId:"api",mdxType:"Tabs"},(0,l.kt)(i.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren template view cpu\n"))),(0,l.kt)(i.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,l.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",h,"/",b,"/templates/cpu"))),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Fetching by Tags")),(0,l.kt)("p",null,"Here is an example to fetch a templates matching the tag."),(0,l.kt)(r.Z,{groupId:"api",mdxType:"Tabs"},(0,l.kt)(i.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren template list --tag firehose\n"))),(0,l.kt)(i.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,l.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request GET\n  --url ",h,"/",b,"/templates?tag=firehose"))),(0,l.kt)("h3",{id:"deleting-a-template"},"Deleting a template"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-text"},"DELETE /v1beta1/templates/cpu HTTP/1.1\nHost: localhost:3000\n")),(0,l.kt)(r.Z,{groupId:"api",mdxType:"Tabs"},(0,l.kt)(i.Z,{value:"cli",label:"CLI",default:!0,mdxType:"TabItem"},(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"$ siren template delete cpu\n"))),(0,l.kt)(i.Z,{value:"http",label:"HTTP",mdxType:"TabItem"},(0,l.kt)(o.Z,{className:"language-bash",mdxType:"CodeBlock"},"$ curl --request DELETE\n  --url ",h,"/",b,"/templates/cpu"))),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Note:")),(0,l.kt)("ol",null,(0,l.kt)("li",{parentName:"ol"},"Updating a template via API will not upload the associated rules.")),(0,l.kt)("h2",{id:"cli-interface"},"CLI interface"),(0,l.kt)("p",null,"With CLI, you will need a YAML file in the below specified format to create/update templates. The CLI calls Siren\nservice templates APIs in turn."),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Example template file")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: v2\ntype: template\nname: CPU\nbody:\n  - alert: CPUWarning\n    expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.warning]]\n    for: "[[.for]]"\n    labels:\n      severity: WARNING\n    annotations:\n      description: CPU has been above [[.warning]] for last [[.for]] {{ $labels.host }}\n  - alert: CPUCritical\n    expr: avg by (host) (cpu_usage_user{cpu="cpu-total"}) > [[.critical]]\n    for: "[[.for]]"\n    labels:\n      severity: CRITICAL\n    annotations:\n      description: CPU has been above [[.critical]] for last [[.for]] {{ $labels.host }}\nvariables:\n  - name: for\n    type: string\n    default: 10m\n    description: For eg 5m, 2h; Golang duration format\n  - name: warning\n    type: int\n    default: 80\n  - name: critical\n    type: int\n    default: 90\ntags:\n  - systems\n')),(0,l.kt)("p",null,"In the above example, we are using one template to define rules of two severity labels viz WARNING and CRITICAL. Here we\nhave made 3 templates variables ",(0,l.kt)("inlineCode",{parentName:"p"},"for"),", ",(0,l.kt)("inlineCode",{parentName:"p"},"warning")," and ",(0,l.kt)("inlineCode",{parentName:"p"},"critical")," which denote the appropriate alerting thresholds. They\nwill be given a value while actual rule(alert) creating."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-bash"},"siren template upload cpu_template.yaml\n")),(0,l.kt)("h3",{id:"terminology"},"Terminology"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Term"),(0,l.kt)("th",{parentName:"tr",align:null},"Description"),(0,l.kt)("th",{parentName:"tr",align:null},"Example/Default"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"API Version"),(0,l.kt)("td",{parentName:"tr",align:null},"Which API to use to parse the YAML file"),(0,l.kt)("td",{parentName:"tr",align:null},"v2")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Type"),(0,l.kt)("td",{parentName:"tr",align:null},"Describes the type of object represented by YAML file"),(0,l.kt)("td",{parentName:"tr",align:null},"template")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Name"),(0,l.kt)("td",{parentName:"tr",align:null},"Name of the template"),(0,l.kt)("td",{parentName:"tr",align:null},"CPUHigh")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Body"),(0,l.kt)("td",{parentName:"tr",align:null},"Array of rule body. The body can be templatized in go template format."),(0,l.kt)("td",{parentName:"tr",align:null},"See example file")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Variables"),(0,l.kt)("td",{parentName:"tr",align:null},"Array of variables that were templatized in the body with their data type, default value and description."),(0,l.kt)("td",{parentName:"tr",align:null},"See example file")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Tags"),(0,l.kt)("td",{parentName:"tr",align:null},"Array of resources/applications that can utilize this template"),(0,l.kt)("td",{parentName:"tr",align:null},"VM")))),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Note:")),(0,l.kt)("ol",null,(0,l.kt)("li",{parentName:"ol"},"It's suggested to always provide default value for the templated variables."),(0,l.kt)("li",{parentName:"ol"},"Updating a template used by rules via CLI will update all associated rules.")))}f.isMDXComponent=!0},3618:(e,t,a)=>{"use strict";a.r(t),a.d(t,{default:()=>n});const n={plain:{color:"#F8F8F2",backgroundColor:"#282A36"},styles:[{types:["prolog","constant","builtin"],style:{color:"rgb(189, 147, 249)"}},{types:["inserted","function"],style:{color:"rgb(80, 250, 123)"}},{types:["deleted"],style:{color:"rgb(255, 85, 85)"}},{types:["changed"],style:{color:"rgb(255, 184, 108)"}},{types:["punctuation","symbol"],style:{color:"rgb(248, 248, 242)"}},{types:["string","char","tag","selector"],style:{color:"rgb(255, 121, 198)"}},{types:["keyword","variable"],style:{color:"rgb(189, 147, 249)",fontStyle:"italic"}},{types:["comment"],style:{color:"rgb(98, 114, 164)"}},{types:["attr-name"],style:{color:"rgb(241, 250, 140)"}}]}},7694:(e,t,a)=>{"use strict";a.r(t),a.d(t,{default:()=>n});const n={plain:{color:"#393A34",backgroundColor:"#f6f8fa"},styles:[{types:["comment","prolog","doctype","cdata"],style:{color:"#999988",fontStyle:"italic"}},{types:["namespace"],style:{opacity:.7}},{types:["string","attr-value"],style:{color:"#e3116c"}},{types:["punctuation","operator"],style:{color:"#393A34"}},{types:["entity","url","symbol","number","boolean","variable","constant","property","regex","inserted"],style:{color:"#36acaa"}},{types:["atrule","keyword","attr-name","selector"],style:{color:"#00a4db"}},{types:["function","deleted","tag"],style:{color:"#d73a49"}},{types:["function-variable"],style:{color:"#6f42c1"}},{types:["tag","selector","keyword"],style:{color:"#00009f"}}]}}}]);