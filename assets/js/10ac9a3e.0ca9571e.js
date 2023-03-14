"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[81],{3905:(e,t,n)=>{n.d(t,{Zo:()=>d,kt:()=>f});var i=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,i,r=function(e,t){if(null==e)return{};var n,i,r={},o=Object.keys(e);for(i=0;i<o.length;i++)n=o[i],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(i=0;i<o.length;i++)n=o[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var l=i.createContext({}),c=function(e){var t=i.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},d=function(e){var t=c(e.components);return i.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},p=i.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,l=e.parentName,d=s(e,["components","mdxType","originalType","parentName"]),p=c(n),f=r,g=p["".concat(l,".").concat(f)]||p[f]||u[f]||o;return n?i.createElement(g,a(a({ref:t},d),{},{components:n})):i.createElement(g,a({ref:t},d))}));function f(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,a=new Array(o);a[0]=p;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:r,a[1]=s;for(var c=2;c<o;c++)a[c]=n[c];return i.createElement.apply(null,a)}return i.createElement.apply(null,n)}p.displayName="MDXCreateElement"},4401:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>u,frontMatter:()=>o,metadata:()=>s,toc:()=>c});var i=n(3117),r=(n(7294),n(3905));const o={},a="Use Cases",s={unversionedId:"use_cases",id:"use_cases",title:"Use Cases",description:"As an Incident Management Platform, Siren integrates with several monitoring and alerting providers (CortexMetrics, Prometheus, InfluxDB, etc) and orchestrates alerting rules in a simple DIY configuration. Siren capables to subscribe to alerts and send notifications based on the triggered alerts or sending on-demand notifications to the supported receivers (slack, pagerduty, etc).",source:"@site/docs/use_cases.md",sourceDirName:".",slug:"/use_cases",permalink:"/siren/docs/use_cases",draft:!1,editUrl:"https://github.com/goto/siren/edit/master/docs/docs/use_cases.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Introduction",permalink:"/siren/docs/introduction"},next:{title:"Installation",permalink:"/siren/docs/installation"}},l={},c=[{value:"Alerting",id:"alerting",level:2},{value:"Alerting Rules Orchestration",id:"alerting-rules-orchestration",level:3},{value:"Alerting Rules Templating",id:"alerting-rules-templating",level:3},{value:"Notification",id:"notification",level:2},{value:"Alert Notifications Subscription",id:"alert-notifications-subscription",level:3},{value:"Sending On-demand Notification",id:"sending-on-demand-notification",level:3}],d={toc:c};function u(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,i.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"use-cases"},"Use Cases"),(0,r.kt)("p",null,"As an Incident Management Platform, Siren integrates with several monitoring and alerting providers (CortexMetrics, Prometheus, InfluxDB, etc) and orchestrates alerting rules in a simple DIY configuration. Siren capables to subscribe to alerts and send notifications based on the triggered alerts or sending on-demand notifications to the supported receivers (slack, pagerduty, etc)."),(0,r.kt)("p",null,"This page describes some of Siren use cases and provides some related resources for better understanding. There might be some other use cases not mentioned in this page that are also suitable with Siren."),(0,r.kt)("h2",{id:"alerting"},"Alerting"),(0,r.kt)("h3",{id:"alerting-rules-orchestration"},"Alerting Rules Orchestration"),(0,r.kt)("p",null,"A rule is an expression that should be met, given the metrics, to trigger an alert. Each monitoring & alerting provider has its own way to define alerting rules and it is relatively easy to do so. However it does not give that much flexibility when the users and teams are getting bigger and there is a need to do self-serve alerting rules creation. Siren provides an abstraction on top of that to give more flexibility in creating alerting rules (via API, CLI, or a UI)."),(0,r.kt)("h3",{id:"alerting-rules-templating"},"Alerting Rules Templating"),(0,r.kt)("p",null,"We noticed there are several times when multiple users or teams using the same rules with just different threshold numbers or labels. Creating multiple similar rules for different purposes is not necessary and would give more overhead to maintain. Siren provides ",(0,r.kt)("a",{parentName:"p",href:"/siren/docs/guides/template"},"templating")," feature to templatize rules given some variables so users could reuse the existing templates to define rules."),(0,r.kt)("h2",{id:"notification"},"Notification"),(0,r.kt)("h3",{id:"alert-notifications-subscription"},"Alert Notifications Subscription"),(0,r.kt)("p",null,"Most monitoring and alerting providers have their own feature to notify a specific channel when an alert is triggered. If an organization uses different monitoring and alerting providers, the responsibility to send notification would be passed on to the respective providers. The number of supported notification channels is also vary depends on the provider. This would give limitation if one needs to send a notification to the unsupported channel in a provider. "),(0,r.kt)("p",null,"With Siren, notification responsibility will be unified in Siren. This approach will be more maintainable and easier to audit. Siren handles all alert notification subscriptions where user could define subscriptions and Siren publishes notifications if the labels in subscriptions match with the labels in the triggered alerts."),(0,r.kt)("p",null,"Siren is also designed to be easily extended with a new notification channel as a ",(0,r.kt)("a",{parentName:"p",href:"/siren/docs/extend/adding_new_receiver"},"new receiver plugin")," to support more use cases."),(0,r.kt)("h3",{id:"sending-on-demand-notification"},"Sending On-demand Notification"),(0,r.kt)("p",null,"There is also a case when a non-alert event needs to be sent as notification with a custom payload. Siren could be used to send on-demand ",(0,r.kt)("a",{parentName:"p",href:"/siren/docs/guides/notification"},"notifications")," too. One just need to pick to which receiver to send notifications too or create a new one if it does not exist yet and send a notification to it."))}u.isMDXComponent=!0}}]);