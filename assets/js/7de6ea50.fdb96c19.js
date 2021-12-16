"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[510],{3905:function(e,n,t){t.d(n,{Zo:function(){return s},kt:function(){return d}});var r=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function a(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function c(e,n){if(null==e)return{};var t,r,i=function(e,n){if(null==e)return{};var t,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var u=r.createContext({}),l=function(e){var n=r.useContext(u),t=n;return e&&(t="function"==typeof e?e(n):a(a({},n),e)),t},s=function(e){var n=l(e.components);return r.createElement(u.Provider,{value:n},e.children)},p={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},m=r.forwardRef((function(e,n){var t=e.components,i=e.mdxType,o=e.originalType,u=e.parentName,s=c(e,["components","mdxType","originalType","parentName"]),m=l(t),d=i,f=m["".concat(u,".").concat(d)]||m[d]||p[d]||o;return t?r.createElement(f,a(a({ref:n},s),{},{components:t})):r.createElement(f,a({ref:n},s))}));function d(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var o=t.length,a=new Array(o);a[0]=m;var c={};for(var u in n)hasOwnProperty.call(n,u)&&(c[u]=n[u]);c.originalType=e,c.mdxType="string"==typeof e?e:i,a[1]=c;for(var l=2;l<o;l++)a[l]=t[l];return r.createElement.apply(null,a)}return r.createElement.apply(null,t)}m.displayName="MDXCreateElement"},5025:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return c},contentTitle:function(){return u},metadata:function(){return l},toc:function(){return s},default:function(){return m}});var r=t(7462),i=t(3366),o=(t(7294),t(3905)),a=["components"],c={},u="Monitoring",l={unversionedId:"guides/monitoring",id:"guides/monitoring",isDocsHomePage:!1,title:"Monitoring",description:"Siren comes with New relic monitoring in built, which user can enable from inside the config.yaml",source:"@site/docs/guides/monitoring.md",sourceDirName:"guides",slug:"/guides/monitoring",permalink:"/siren/docs/guides/monitoring",editUrl:"https://github.com/odpf/siren/edit/master/docs/docs/guides/monitoring.md",tags:[],version:"current",lastUpdatedBy:"Ravi Suhag",lastUpdatedAt:1639650958,formattedLastUpdatedAt:"12/16/2021",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Bulk Rule management",permalink:"/siren/docs/guides/bulk_rules"},next:{title:"Deployment",permalink:"/siren/docs/guides/deployment"}},s=[],p={toc:s};function m(e){var n=e.components,t=(0,i.Z)(e,a);return(0,o.kt)("wrapper",(0,r.Z)({},p,t,{components:n,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"monitoring"},"Monitoring"),(0,o.kt)("p",null,"Siren comes with New relic monitoring in built, which user can enable from inside the ",(0,o.kt)("inlineCode",{parentName:"p"},"config.yaml")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-yaml"},"newrelic:\n  enabled: true\n  appname: siren\n  license: ____LICENSE_STRING_OF_40_CHARACTERS_____\n")),(0,o.kt)("p",null,"If the ",(0,o.kt)("inlineCode",{parentName:"p"},"enabled")," is set to true, with correct ",(0,o.kt)("inlineCode",{parentName:"p"},"license")," key, you will be able to see the API metrics on your New relic\ndashboard."))}m.isMDXComponent=!0}}]);