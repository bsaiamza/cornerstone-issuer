"use strict";(self.webpackChunkcornerstone_issuer=self.webpackChunkcornerstone_issuer||[]).push([[562],{9624:function(n,e,t){var i=t(4567),o=t(184);e.Z=function(n){var e=n.align,t=n.children,r=n.color,c=n.sx,s=n.variant;return(0,o.jsx)(i.Z,{align:e,color:r||"#777",sx:c,variant:s,children:t})}},2562:function(n,e,t){t.r(e);var i=t(4165),o=t(5861),r=t(885),c=t(2791),s=t(4569),a=t.n(s),_=t(9098),u=t(3439),d=t(5985),S=t(9624),l=t(184);e.default=function(){var n=(0,c.useState)([]),e=(0,r.Z)(n,2),t=e[0],s=e[1];(0,c.useEffect)((function(){var n="/api/v1/cornerstone/issuer/connections";({NODE_ENV:"production",PUBLIC_URL:"",WDS_SOCKET_HOST:void 0,WDS_SOCKET_PATH:void 0,WDS_SOCKET_PORT:void 0,FAST_REFRESH:!0}).API_BASE_URL&&(n={NODE_ENV:"production",PUBLIC_URL:"",WDS_SOCKET_HOST:void 0,WDS_SOCKET_PATH:void 0,WDS_SOCKET_PORT:void 0,FAST_REFRESH:!0}.API_BASE_URL+"/cornerstone/issuer/connections"),a().get(n+"?state=active").then((function(n){s(n.data)})).catch((function(n){return console.log(n)}))}),[]);var v=function(){var n=(0,o.Z)((0,i.Z)().mark((function n(){var e;return(0,i.Z)().wrap((function(n){for(;;)switch(n.prev=n.next){case 0:e="/api/v1/cornerstone/issuer/connections",{NODE_ENV:"production",PUBLIC_URL:"",WDS_SOCKET_HOST:void 0,WDS_SOCKET_PATH:void 0,WDS_SOCKET_PORT:void 0,FAST_REFRESH:!0}.API_BASE_URL&&(e={NODE_ENV:"production",PUBLIC_URL:"",WDS_SOCKET_HOST:void 0,WDS_SOCKET_PATH:void 0,WDS_SOCKET_PORT:void 0,FAST_REFRESH:!0}.API_BASE_URL+"/cornerstone/issuer/connections"),n.next=7;break;case 5:n.next=9;break;case 7:return n.next=9,d.Am.promise(a().get(e+"?state=active").then((function(n){s(n.data),d.Am.success("Refreshed connections!")})).catch((function(n){return console.log(n)})),{pending:"Refreshing..."});case 9:case"end":return n.stop()}}),n)})));return function(){return n.apply(this,arguments)}}();return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)("div",{style:{margin:"2rem"},children:(0,l.jsx)(_.ZP,{style:{padding:"1rem"},title:(0,l.jsx)(S.Z,{variant:"h5",sx:{textDecoration:"underline"},children:"Connections"}),data:t,columns:[{title:(0,l.jsx)(S.Z,{variant:"h6",children:"Name"}),field:"their_label"},{title:(0,l.jsx)(S.Z,{variant:"h6",children:"Connected On"}),field:"created_at",type:"datetime"},{title:(0,l.jsx)(S.Z,{variant:"h6",children:"Connection ID"}),field:"connection_id"},{title:(0,l.jsx)(S.Z,{variant:"h6",children:"Connection State"}),field:"state"}],actions:[{icon:function(){return(0,l.jsx)(u.Z,{})},tooltip:"Refresh connections",isFreeAction:!0,onClick:function(){return v()}}],options:{actionsColumnIndex:-1}})}),(0,l.jsx)("div",{style:{marginBottom:"2rem"}})]})}}}]);
//# sourceMappingURL=562.9744ded7.chunk.js.map