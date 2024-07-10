"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[2545],{67749:(e,r,s)=>{s.r(r),s.d(r,{assets:()=>l,contentTitle:()=>t,default:()=>h,frontMatter:()=>o,metadata:()=>a,toc:()=>u});var n=s(58168),i=(s(96540),s(15680));const o={},t="Workers",a={unversionedId:"guides/workers",id:"guides/workers",title:"Workers",description:"Siren has a notification features that utilizes queue to publish notification messages. More concept about notification could be found here. The architecture requires a detached worker running asynchronously and polling queue periodically to dequeue notification messages and publish them. By default, Siren server run this asynchronous worker inside it. However it is also possible to run the worker as a different process. Currently there are two possible workers to run",source:"@site/docs/guides/workers.md",sourceDirName:"guides",slug:"/guides/workers",permalink:"/siren/docs/guides/workers",draft:!1,editUrl:"https://github.com/goto/siren/edit/master/docs/docs/guides/workers.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Notification",permalink:"/siren/docs/guides/notification"},next:{title:"Job",permalink:"/siren/docs/guides/job"}},l={},u=[{value:"Running Workers as a Different Process",id:"running-workers-as-a-different-process",level:2}],d={toc:u},c="wrapper";function h(e){let{components:r,...s}=e;return(0,i.yg)(c,(0,n.A)({},d,s,{components:r,mdxType:"MDXLayout"}),(0,i.yg)("h1",{id:"workers"},"Workers"),(0,i.yg)("p",null,"Siren has a notification features that utilizes queue to publish notification messages. More concept about notification could be found ",(0,i.yg)("a",{parentName:"p",href:"/siren/docs/concepts/notification"},"here"),". The architecture requires a detached worker running asynchronously and polling queue periodically to dequeue notification messages and publish them. By default, Siren server run this asynchronous worker inside it. However it is also possible to run the worker as a different process. Currently there are two possible workers to run"),(0,i.yg)("ol",null,(0,i.yg)("li",{parentName:"ol"},(0,i.yg)("strong",{parentName:"li"},"Notification message handler:")," this worker periodically poll and dequeue messages from queue, process the messages, and then publish notification messages to the specified receivers. If there is a failure, this handler enqueues the failed messages to the dlq."),(0,i.yg)("li",{parentName:"ol"},(0,i.yg)("strong",{parentName:"li"},"Notification dlq handler:")," this worker periodically poll and dequeue messages from dead-letter-queue, process the messages, and then publish notification messages to the specified receivers. If there is a failure, this handler enqueues the failed messages back to the dlq.")),(0,i.yg)("h2",{id:"running-workers-as-a-different-process"},"Running Workers as a Different Process"),(0,i.yg)("p",null,"Siren has a command to start workers. Workers use the same config like server does."),(0,i.yg)("pre",null,(0,i.yg)("code",{parentName:"pre",className:"language-bash"},"Command to start a siren worker.\n\nUsage\n  siren worker start <command> [flags]\n\nCore commands\n  notification_dlq_handler    A notification dlq handler\n  notification_handler        A notification handler\n\nInherited flags\n  --help   Show help for command\n\nExamples\n  $ siren worker start notification_handler\n  $ siren server start notification_handler -c ./config.yaml\n")),(0,i.yg)("p",null,"Starting up a worker could be done by executing."),(0,i.yg)("pre",null,(0,i.yg)("code",{parentName:"pre",className:"language-bash"},"$ siren worker start notification_handler\n")))}h.isMDXComponent=!0}}]);