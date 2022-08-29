"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[700],{3905:function(e,t,n){n.d(t,{Zo:function(){return s},kt:function(){return c}});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},l=Object.keys(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var p=a.createContext({}),u=function(e){var t=a.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},s=function(e){var t=u(e.components);return a.createElement(p.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,l=e.originalType,p=e.parentName,s=o(e,["components","mdxType","originalType","parentName"]),m=u(n),c=r,g=m["".concat(p,".").concat(c)]||m[c]||d[c]||l;return n?a.createElement(g,i(i({ref:t},s),{},{components:n})):a.createElement(g,i({ref:t},s))}));function c(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=n.length,i=new Array(l);i[0]=m;var o={};for(var p in t)hasOwnProperty.call(t,p)&&(o[p]=t[p]);o.originalType=e,o.mdxType="string"==typeof e?e:r,i[1]=o;for(var u=2;u<l;u++)i[u]=n[u];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},9157:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return o},contentTitle:function(){return p},metadata:function(){return u},toc:function(){return s},default:function(){return m}});var a=n(7462),r=n(3366),l=(n(7294),n(3905)),i=["components"],o={},p="Rules",u={unversionedId:"guides/rules",id:"guides/rules",isDocsHomePage:!1,title:"Rules",description:"Siren rules are creating from predefined templates by providing values of the variables of the template.",source:"@site/docs/guides/rules.md",sourceDirName:"guides",slug:"/guides/rules",permalink:"/siren/docs/guides/rules",editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/rules.md",tags:[],version:"current",lastUpdatedBy:"Ravi Suhag",lastUpdatedAt:1661746387,formattedLastUpdatedAt:"8/29/2022",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Subscriptions",permalink:"/siren/docs/guides/subscriptions"},next:{title:"Templates",permalink:"/siren/docs/guides/templates"}},s=[{value:"API interface",id:"api-interface",children:[{value:"Terminology of the request body",id:"terminology-of-the-request-body",children:[]}]},{value:"CLI Interface",id:"cli-interface",children:[{value:"Terminology",id:"terminology",children:[]}]}],d={toc:s};function m(e){var t=e.components,n=(0,r.Z)(e,i);return(0,l.kt)("wrapper",(0,a.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h1",{id:"rules"},"Rules"),(0,l.kt)("p",null,"Siren rules are creating from predefined ",(0,l.kt)("a",{parentName:"p",href:"/siren/docs/guides/templates"},"templates")," by providing values of the variables of the template."),(0,l.kt)("p",null,"One can create templates using either HTTP APIs or CLI."),(0,l.kt)("h2",{id:"api-interface"},"API interface"),(0,l.kt)("p",null,"A rule is uniquely identified with the combination of provider's namespace(uniquely identifies which provider and\nnamespace), template name, optional namespace, optional group name."),(0,l.kt)("p",null,"One can choose any namespace and group name. In cortex terminology, namespace is a collection of groups. Groups can have\none or more rules."),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Creating/Updating a rule")),(0,l.kt)("p",null,"The below snippet describes an example of rule creation/update. Same API can be used to enable or disable alerts."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-text"},'PUT /v1beta1/rules HTTP/1.1\nHost: localhost:3000\nContent-Type: application/json\nContent-Length: 298\n\n{\n  "namespace": "odpf",\n  "group_name": "CPUHigh",\n  "template": "CPU",\n  "providerNamespace": "3"\n  "variables": [\n    {\n      "name": "for",\n      "value": "15m",\n      "type": "string"\n    },\n     {\n      "name": "team",\n      "value": "odpf",\n      "type": "string"\n    }\n  ],\n  "enabled": true,\n}\n')),(0,l.kt)("p",null,'Here we are using CPU template and providing value for few variables("for", "team"). In case some variables value is not\nprovided default will be picked from the template\'s definition.'),(0,l.kt)("h3",{id:"terminology-of-the-request-body"},"Terminology of the request body"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Term"),(0,l.kt)("th",{parentName:"tr",align:null},"Description"),(0,l.kt)("th",{parentName:"tr",align:null},"Example/Default"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Namespace"),(0,l.kt)("td",{parentName:"tr",align:null},"Corresponds to Cortex namespace in which rule will be created"),(0,l.kt)("td",{parentName:"tr",align:null},"kafka")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Group Name"),(0,l.kt)("td",{parentName:"tr",align:null},"Corresponds to Cortex group name in which rule will be created"),(0,l.kt)("td",{parentName:"tr",align:null},"CPUHigh")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"providerNamespace"),(0,l.kt)("td",{parentName:"tr",align:null},"Corresponds to a tenant in a provider"),(0,l.kt)("td",{parentName:"tr",align:null},"4")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Template"),(0,l.kt)("td",{parentName:"tr",align:null},"what template is used to create the rule"),(0,l.kt)("td",{parentName:"tr",align:null},"CPU")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Variables"),(0,l.kt)("td",{parentName:"tr",align:null},"Value of variables defined inside the template"),(0,l.kt)("td",{parentName:"tr",align:null},"See example above")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Enabled"),(0,l.kt)("td",{parentName:"tr",align:null},"boolean describing if the rule is enabled or not"),(0,l.kt)("td",{parentName:"tr",align:null},"true")))),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Fetching rules")),(0,l.kt)("p",null,"Rules can be fetched and filtered with multiple parameters. An example of all filters is described below."),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-text"},"GET /v1beta1/rules?namespace=foo&providerNamespace=4&group_name=CPUHigh&template=CPU HTTP/1.1\nHost: localhost:3000\n")),(0,l.kt)("h2",{id:"cli-interface"},"CLI Interface"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-text"},'Work with rules.\n\nrules are used for alerting within a provider.\n\nUsage:\n  siren rule [command]\n\nAliases:\n  rule, rules\n\nAvailable Commands:\n  edit        Edit a rule\n  list        List rules\n  upload      Upload Rules YAML file\n\nFlags:\n  -h, --help   help for rule\n\nUse "siren rule [command] --help" for more information about a command.\n')),(0,l.kt)("p",null,"With CLI, you will need a YAML file in the below specified format to create/edit rules.\n",(0,l.kt)("strong",{parentName:"p"},"Example rule file")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: v2\ntype: rule\nnamespace: demo\nprovider: localhost-cortex\nproviderNamespace: test\nrules:\n  TestGroup:\n    template: CPU\n    status: enabled\n    variables:\n      - name: for\n        value: 15m\n      - name: warning\n        value: 185\n      - name: critical\n        value: 195\n")),(0,l.kt)("p",null,"In the above example, we are creating rules\ninside ",(0,l.kt)("inlineCode",{parentName:"p"},"demo")," ",(0,l.kt)("a",{parentName:"p",href:"https://cortexmetrics.io/docs/api/#get-rule-groups-by-namespace"},"namespace"),"\nunder ",(0,l.kt)("inlineCode",{parentName:"p"},"test")," ",(0,l.kt)("a",{parentName:"p",href:"https://cortexmetrics.io/docs/architecture/#the-role-of-prometheus"},"tenant")," of ",(0,l.kt)("inlineCode",{parentName:"p"},"localhost-cortex"),"\nprovider."),(0,l.kt)("p",null,"The rules array defines actual rules defined over the templates. Here ",(0,l.kt)("inlineCode",{parentName:"p"},"TestGroup")," is the name of the group which will be\ncreated/updated with the rule defined by ",(0,l.kt)("inlineCode",{parentName:"p"},"CPU")," template. The example shows the value of variables provided in creating\nrules(alert)."),(0,l.kt)("p",null,(0,l.kt)("strong",{parentName:"p"},"Example upload command")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-shell"},"go run main.go rule upload cpu_rule.yaml\n")),(0,l.kt)("p",null,"The yaml file can be edited and re-uploaded to edit the rule thresholds."),(0,l.kt)("h3",{id:"terminology"},"Terminology"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Term"),(0,l.kt)("th",{parentName:"tr",align:null},"Description"),(0,l.kt)("th",{parentName:"tr",align:null},"Example/Default"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"API Version"),(0,l.kt)("td",{parentName:"tr",align:null},"Which API to use to parse the YAML file"),(0,l.kt)("td",{parentName:"tr",align:null},"v2")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Type"),(0,l.kt)("td",{parentName:"tr",align:null},"Describes the type of object represented by YAML file"),(0,l.kt)("td",{parentName:"tr",align:null},"rule")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Namespace"),(0,l.kt)("td",{parentName:"tr",align:null},"Corresponds to Cortex namespace in which rule will be created"),(0,l.kt)("td",{parentName:"tr",align:null},"kafka")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Entity"),(0,l.kt)("td",{parentName:"tr",align:null},"Corresponds to tenant name in cortex"),(0,l.kt)("td",{parentName:"tr",align:null},"odpf")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Rules"),(0,l.kt)("td",{parentName:"tr",align:null},"Map of GroupNames describing what template is used in a particular group"),(0,l.kt)("td",{parentName:"tr",align:null},"See example file")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"Variables"),(0,l.kt)("td",{parentName:"tr",align:null},"Value of variables defined inside the template"),(0,l.kt)("td",{parentName:"tr",align:null},"See example above")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"provider"),(0,l.kt)("td",{parentName:"tr",align:null},"URN of monitoring provider to be used"),(0,l.kt)("td",{parentName:"tr",align:null},"localhost-cortex")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"providerNamespace"),(0,l.kt)("td",{parentName:"tr",align:null},"URN of tenant to choose inside the monitoring provider"),(0,l.kt)("td",{parentName:"tr",align:null},"test")))))}m.isMDXComponent=!0}}]);