"use strict";(self.webpackChunkcornerstone_issuer=self.webpackChunkcornerstone_issuer||[]).push([[716],{9055:function(e,n,t){var r=t(4554),i=t(184);n.Z=function(e){var n=e.alt,t=e.children,a=e.component,o=e.src,s=e.sx;return(0,i.jsx)(r.Z,{alt:n,component:a,src:o,sx:s,children:t})}},816:function(e,n,t){var r=t(439),i=t(184);n.Z=function(e){var n=e.children,t=e.color,a=e.disabled,o=e.endIcon,s=e.onClick,c=e.size,d=e.startIcon,l=e.sx,u=e.type,v=e.variant;return(0,i.jsx)(r.Z,{color:t,disabled:a,endIcon:o,onClick:s,size:c,sx:l,startIcon:d,type:u,variant:v,children:n})}},656:function(e,n,t){var r=t(7621),i=t(9504),a=t(184);n.Z=function(e){var n=e.children,t=e.sx;return(0,a.jsx)(r.Z,{sx:t,elevation:3,children:(0,a.jsx)(i.Z,{children:n})})}},5589:function(e,n,t){var r=t(2506),i=t(184);n.Z=function(e){var n=e.children;return(0,i.jsx)(r.l0,{children:n})}},8425:function(e,n,t){var r=t(2506),i=t(184);n.Z=function(e){var n=e.children,t=e.initialValues,a=e.onSubmit;return(0,i.jsx)(r.J9,{initialValues:t,onSubmit:a,children:n})}},6321:function(e,n,t){var r=t(1889),i=t(184);n.Z=function(e){var n=e.children,t=e.container,a=e.item,o=e.justify,s=e.md,c=e.spacing,d=e.xs;return(0,i.jsx)(r.ZP,{container:t,item:a,justify:o,md:s,spacing:c,xs:d,children:n})}},3434:function(e,n,t){var r=t(5021),i=t(184);n.Z=function(e){var n=e.children,t=e.disablePadding,a=e.onClick,o=e.sx;return(0,i.jsx)(r.ZP,{disablePadding:t,onClick:a,sx:o,children:n})}},7774:function(e,n,t){var r=t(9900),i=(t(2791),t(184));n.Z=function(e){var n=e.color,t=e.primary,a=e.sx;return(0,i.jsx)(r.Z,{color:n,primary:t,sx:{sx:a}})}},6484:function(e,n,t){var r=t(1413),i=t(5987),a=t(3896),o=t(184),s=["label"];n.Z=function(e){var n=e.label,t=(0,i.Z)(e,s);return(0,o.jsx)(a.Z,(0,r.Z)({label:n},t))}},4373:function(e,n,t){t.d(n,{P:function(){return l},x:function(){return d}});var r=t(1413),i=t(5987),a=t(9055),o=t(9624),s=t(184),c=["children","value","index"],d=function(e){var n=e.children,t=e.value,d=e.index,l=(0,i.Z)(e,c);return(0,s.jsx)("div",(0,r.Z)((0,r.Z)({role:"tabpanel",hidden:t!==d,id:"custom-tabpanel-".concat(d),"aria-labelledby":"custom-tab-".concat(d)},l),{},{children:t===d&&(0,s.jsx)(a.Z,{sx:{p:3},children:(0,s.jsx)(o.Z,{children:n})})}))},l=function(e){return{id:"custom-tab-".concat(e),"aria-controls":"custom-tabpanel-".concat(e)}}},5192:function(e,n,t){var r=t(2101),i=t(184);n.Z=function(e){var n=e.ariaLabel,t=e.children,a=e.onChange,o=e.value;return(0,i.jsx)(r.Z,{"aria-label":n,onChange:a,value:o,children:t})}},4695:function(e,n,t){var r=t(3006),i=t(184);n.Z=function(e){var n=e.disabled,t=e.focused,a=e.fullWidth,o=e.id,s=e.label,c=e.name,d=e.onChange,l=e.required,u=e.sx,v=e.type,p=e.value;return(0,i.jsx)(r.Z,{disabled:n,focused:t,fullWidth:a,id:o,label:s,name:c,onChange:d,required:l,sx:u,type:v,value:p})}},9624:function(e,n,t){var r=t(4567),i=t(184);n.Z=function(e){var n=e.align,t=e.children,a=e.color,o=e.sx,s=e.variant;return(0,i.jsx)(r.Z,{align:n,color:a||"#777",sx:o,variant:s,children:t})}},3716:function(e,n,t){t.r(n),t.d(n,{default:function(){return D}});var r=t(1413),i=t(885),a=t(2791),o=t(9055),s=t(5192),c=t(6484),d=t(4373),l=t(4569),u=t.n(l),v=t(6960),p=t(816),m=t(656),Z=t(6321),f=t(3434),h=t(7774),x=t(184),g=function(){var e=(0,a.useState)([]),n=(0,i.Z)(e,2),t=n[0],r=n[1],o=(0,a.useState)([]),s=(0,i.Z)(o,2),c=s[0],d=s[1],l=(0,a.useState)([]),g=(0,i.Z)(l,2),b=g[0],j=g[1];(0,a.useEffect)((function(){u().get("/api/v1/cornerstone/issuer/did").then((function(e){r(e.data)})).catch((function(e){return console.log(e)})),"undefined"!==t.did&&u().get("/api/v1/cornerstone/issuer/definitions?issuer_did="+t.did).then((function(e){d(e.data.credential_definition_ids)})).catch((function(e){return console.log(e)}))}),[t.did]);return(0,x.jsxs)(Z.Z,{container:!0,spacing:2,children:[(0,x.jsx)(Z.Z,{item:!0,xs:12,md:4,children:c?0===c.length?(0,x.jsx)(m.Z,{sx:{m:"1rem"},children:"No credential definitions available!"}):c.map((function(e,n){return(0,x.jsx)(m.Z,{sx:{m:"1rem",overflowWrap:"break-word"},children:(0,x.jsxs)(f.Z,{children:["Credential Definition ID:",(0,x.jsx)(p.Z,{onClick:function(){return n=e,void v.Am.promise(u().get("/api/v1/cornerstone/issuer/definition?cred_def_id="+n).then((function(e){j(e.data),v.Am.success("Fetched credential definition!")})).catch((function(e){v.Am.error(e.response.data.msg)})),{pending:"Fetching..."});var n},sx:{color:"#0645AD !important"},children:(0,x.jsx)(h.Z,{primary:e})})]})},n)})):(0,x.jsx)(m.Z,{sx:{m:"1rem"},children:"No credential definitions available!"})}),(0,x.jsx)(Z.Z,{item:!0,xs:12,md:8,children:(0,x.jsxs)(m.Z,{sx:{m:"1rem"},children:[(0,x.jsxs)(f.Z,{children:["Credential Definition ID:"," ",(0,x.jsx)(h.Z,{primary:b.id})]}),(0,x.jsxs)(f.Z,{children:["Schema ID: ",(0,x.jsx)(h.Z,{primary:b.schemaId})]}),(0,x.jsxs)(f.Z,{children:["Type: ",(0,x.jsx)(h.Z,{primary:b.type})]}),(0,x.jsxs)(f.Z,{children:["Tag:",(0,x.jsx)(h.Z,{primary:b.tag})]})]})})]})},b=t(4165),j=t(5861),y=t(5363),C=t(4932),_=t(8524),w=t(2626),I=t(8425),S=t(5589),k=t(4695),M=function(){var e=(0,a.useState)([]),n=(0,i.Z)(e,2),t=n[0],r=n[1],o=(0,a.useState)(!1),s=(0,i.Z)(o,2),c=s[0],d=s[1],l=function(){var e=(0,j.Z)((0,b.Z)().mark((function e(n){return(0,b.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return d(!0),e.next=3,v.Am.promise(u().post("/api/v1/cornerstone/issuer/definition/create",n).then((function(e){console.log(e.data),r(e.data),v.Am.success("Created Credential Definition!")})).catch((function(e){v.Am.error(e.response.data.msg)})),{pending:"Creating..."});case 3:d(!1);case 4:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}();return(0,x.jsx)(m.Z,{children:(0,x.jsxs)(Z.Z,{container:!0,spacing:2,children:[(0,x.jsxs)(Z.Z,{item:!0,xs:12,md:6,children:[(0,x.jsx)("h2",{children:"Add New Credential Definition"}),(0,x.jsx)(I.Z,{initialValues:{schema_id:"",support_revocation:!1,revocation_registry_size:0,tag:""},onSubmit:function(e,n){var t=n.resetForm;l(e),t()},children:function(e){var n=e.values,t=e.handleChange;return(0,x.jsxs)(S.Z,{children:[(0,x.jsx)("div",{children:(0,x.jsx)(k.Z,{id:"schema_id",name:"schema_id",value:n.schema_id,onChange:t,label:"Schema ID",sx:{m:"1rem"},required:!0})}),(0,x.jsx)("div",{children:(0,x.jsxs)(y.Z,{sx:{width:"16.5rem"},children:[(0,x.jsx)(C.Z,{id:"support_revocation_label",sx:{margin:"1rem 0 0 1rem"},children:"Support Revocation"}),(0,x.jsxs)(_.Z,{labelId:"support_revocation_label",id:"support_revocation",name:"support_revocation",value:n.support_revocation,label:"Support Revocation",onChange:t,sx:{m:"1rem"},children:[(0,x.jsx)(w.Z,{value:!1,children:"False"}),(0,x.jsx)(w.Z,{value:!0,children:"True"})]})]})}),(0,x.jsx)("div",{children:(0,x.jsx)(k.Z,{id:"revocation_registry_size",name:"revocation_registry_size",type:"number",value:n.revocation_registry_size,onChange:t,label:"Revocation Registry Size",sx:{m:"1rem"}})}),(0,x.jsx)("div",{children:(0,x.jsx)(k.Z,{id:"tag",name:"tag",value:n.tag,onChange:t,label:"Tag",sx:{m:"1rem"},required:!0})}),(0,x.jsx)("div",{children:(0,x.jsx)(p.Z,{variant:"contained",type:"submit",sx:{color:"#fff",m:"1rem"},disabled:c,children:"Create"})})]})}})]}),(0,x.jsx)(Z.Z,{item:!0,xs:12,md:6,children:t?(0,x.jsx)("div",{style:{padding:"2rem"},children:(0,x.jsx)(f.Z,{children:(0,x.jsx)(h.Z,{primary:"Credential Definition ID: "+t.credential_definition_id})})}):""})]})})},D=function(){var e=(0,a.useState)(0),n=(0,i.Z)(e,2),t=n[0],l=n[1];return(0,x.jsxs)("div",{style:{margin:"1rem"},children:[(0,x.jsx)(o.Z,{children:(0,x.jsxs)(s.Z,{value:t,onChange:function(e,n){l(n)},ariaLabel:"Definition Tabs",children:[(0,x.jsx)(c.Z,(0,r.Z)({label:"All"},(0,d.P)(0))),(0,x.jsx)(c.Z,(0,r.Z)({label:"New"},(0,d.P)(1)))]})}),(0,x.jsx)(d.x,{value:t,index:0,children:(0,x.jsx)(g,{})}),(0,x.jsx)(d.x,{value:t,index:1,children:(0,x.jsx)(M,{})})]})}},7621:function(e,n,t){t.d(n,{Z:function(){return f}});var r=t(7462),i=t(3366),a=t(2791),o=t(8182),s=t(4419),c=t(7630),d=t(3736),l=t(4841),u=t(1217);function v(e){return(0,u.Z)("MuiCard",e)}(0,t(5878).Z)("MuiCard",["root"]);var p=t(184),m=["className","raised"],Z=(0,c.ZP)(l.Z,{name:"MuiCard",slot:"Root",overridesResolver:function(e,n){return n.root}})((function(){return{overflow:"hidden"}})),f=a.forwardRef((function(e,n){var t=(0,d.Z)({props:e,name:"MuiCard"}),a=t.className,c=t.raised,l=void 0!==c&&c,u=(0,i.Z)(t,m),f=(0,r.Z)({},t,{raised:l}),h=function(e){var n=e.classes;return(0,s.Z)({root:["root"]},v,n)}(f);return(0,p.jsx)(Z,(0,r.Z)({className:(0,o.Z)(h.root,a),elevation:l?8:void 0,ref:n,ownerState:f},u))}))},9504:function(e,n,t){t.d(n,{Z:function(){return Z}});var r=t(7462),i=t(3366),a=t(2791),o=t(8182),s=t(4419),c=t(7630),d=t(3736),l=t(1217);function u(e){return(0,l.Z)("MuiCardContent",e)}(0,t(5878).Z)("MuiCardContent",["root"]);var v=t(184),p=["className","component"],m=(0,c.ZP)("div",{name:"MuiCardContent",slot:"Root",overridesResolver:function(e,n){return n.root}})((function(){return{padding:16,"&:last-child":{paddingBottom:24}}})),Z=a.forwardRef((function(e,n){var t=(0,d.Z)({props:e,name:"MuiCardContent"}),a=t.className,c=t.component,l=void 0===c?"div":c,Z=(0,i.Z)(t,p),f=(0,r.Z)({},t,{component:l}),h=function(e){var n=e.classes;return(0,s.Z)({root:["root"]},u,n)}(f);return(0,v.jsx)(m,(0,r.Z)({as:l,className:(0,o.Z)(h.root,a),ownerState:f,ref:n},Z))}))},133:function(e,n,t){t.d(n,{V:function(){return i}});var r=t(1217);function i(e){return(0,r.Z)("MuiDivider",e)}var a=(0,t(5878).Z)("MuiDivider",["root","absolute","fullWidth","inset","middle","flexItem","light","vertical","withChildren","withChildrenVertical","textAlignRight","textAlignLeft","wrapper","wrapperVertical"]);n.Z=a},6014:function(e,n,t){t.d(n,{f:function(){return i}});var r=t(1217);function i(e){return(0,r.Z)("MuiListItemIcon",e)}var a=(0,t(5878).Z)("MuiListItemIcon",["root","alignItemsFlexStart"]);n.Z=a},2626:function(e,n,t){var r=t(4942),i=t(3366),a=t(7462),o=t(2791),s=t(8182),c=t(4419),d=t(2065),l=t(7630),u=t(3736),v=t(6199),p=t(5080),m=t(162),Z=t(2071),f=t(133),h=t(6014),x=t(9849),g=t(1498),b=t(184),j=["autoFocus","component","dense","divider","disableGutters","focusVisibleClassName","role","tabIndex"],y=(0,l.ZP)(p.Z,{shouldForwardProp:function(e){return(0,l.FO)(e)||"classes"===e},name:"MuiMenuItem",slot:"Root",overridesResolver:function(e,n){var t=e.ownerState;return[n.root,t.dense&&n.dense,t.divider&&n.divider,!t.disableGutters&&n.gutters]}})((function(e){var n,t=e.theme,i=e.ownerState;return(0,a.Z)({},t.typography.body1,{display:"flex",justifyContent:"flex-start",alignItems:"center",position:"relative",textDecoration:"none",minHeight:48,paddingTop:6,paddingBottom:6,boxSizing:"border-box",whiteSpace:"nowrap"},!i.disableGutters&&{paddingLeft:16,paddingRight:16},i.divider&&{borderBottom:"1px solid ".concat((t.vars||t).palette.divider),backgroundClip:"padding-box"},(n={"&:hover":{textDecoration:"none",backgroundColor:(t.vars||t).palette.action.hover,"@media (hover: none)":{backgroundColor:"transparent"}}},(0,r.Z)(n,"&.".concat(g.Z.selected),(0,r.Z)({backgroundColor:t.vars?"rgba(".concat(t.vars.palette.primary.mainChannel," / ").concat(t.vars.palette.action.selectedOpacity,")"):(0,d.Fq)(t.palette.primary.main,t.palette.action.selectedOpacity)},"&.".concat(g.Z.focusVisible),{backgroundColor:t.vars?"rgba(".concat(t.vars.palette.primary.mainChannel," / calc(").concat(t.vars.palette.action.selectedOpacity," + ").concat(t.vars.palette.action.focusOpacity,"))"):(0,d.Fq)(t.palette.primary.main,t.palette.action.selectedOpacity+t.palette.action.focusOpacity)})),(0,r.Z)(n,"&.".concat(g.Z.selected,":hover"),{backgroundColor:t.vars?"rgba(".concat(t.vars.palette.primary.mainChannel," / calc(").concat(t.vars.palette.action.selectedOpacity," + ").concat(t.vars.palette.action.hoverOpacity,"))"):(0,d.Fq)(t.palette.primary.main,t.palette.action.selectedOpacity+t.palette.action.hoverOpacity),"@media (hover: none)":{backgroundColor:t.vars?"rgba(".concat(t.vars.palette.primary.mainChannel," / ").concat(t.vars.palette.action.selectedOpacity,")"):(0,d.Fq)(t.palette.primary.main,t.palette.action.selectedOpacity)}}),(0,r.Z)(n,"&.".concat(g.Z.focusVisible),{backgroundColor:(t.vars||t).palette.action.focus}),(0,r.Z)(n,"&.".concat(g.Z.disabled),{opacity:(t.vars||t).palette.action.disabledOpacity}),(0,r.Z)(n,"& + .".concat(f.Z.root),{marginTop:t.spacing(1),marginBottom:t.spacing(1)}),(0,r.Z)(n,"& + .".concat(f.Z.inset),{marginLeft:52}),(0,r.Z)(n,"& .".concat(x.Z.root),{marginTop:0,marginBottom:0}),(0,r.Z)(n,"& .".concat(x.Z.inset),{paddingLeft:36}),(0,r.Z)(n,"& .".concat(h.Z.root),{minWidth:36}),n),!i.dense&&(0,r.Z)({},t.breakpoints.up("sm"),{minHeight:"auto"}),i.dense&&(0,a.Z)({minHeight:32,paddingTop:4,paddingBottom:4},t.typography.body2,(0,r.Z)({},"& .".concat(h.Z.root," svg"),{fontSize:"1.25rem"})))})),C=o.forwardRef((function(e,n){var t=(0,u.Z)({props:e,name:"MuiMenuItem"}),r=t.autoFocus,d=void 0!==r&&r,l=t.component,p=void 0===l?"li":l,f=t.dense,h=void 0!==f&&f,x=t.divider,C=void 0!==x&&x,_=t.disableGutters,w=void 0!==_&&_,I=t.focusVisibleClassName,S=t.role,k=void 0===S?"menuitem":S,M=t.tabIndex,D=(0,i.Z)(t,j),R=o.useContext(v.Z),O={dense:h||R.dense||!1,disableGutters:w},F=o.useRef(null);(0,m.Z)((function(){d&&F.current&&F.current.focus()}),[d]);var N,V=(0,a.Z)({},t,{dense:O.dense,divider:C,disableGutters:w}),P=function(e){var n=e.disabled,t=e.dense,r=e.divider,i=e.disableGutters,o=e.selected,s=e.classes,d={root:["root",t&&"dense",n&&"disabled",!i&&"gutters",r&&"divider",o&&"selected"]},l=(0,c.Z)(d,g.K,s);return(0,a.Z)({},s,l)}(t),A=(0,Z.Z)(F,n);return t.disabled||(N=void 0!==M?M:-1),(0,b.jsx)(v.Z.Provider,{value:O,children:(0,b.jsx)(y,(0,a.Z)({ref:A,role:k,tabIndex:N,component:p,focusVisibleClassName:(0,s.Z)(P.focusVisible,I)},D,{ownerState:V,classes:P}))})}));n.Z=C},1498:function(e,n,t){t.d(n,{K:function(){return i}});var r=t(1217);function i(e){return(0,r.Z)("MuiMenuItem",e)}var a=(0,t(5878).Z)("MuiMenuItem",["root","focusVisible","dense","disabled","divider","gutters","selected"]);n.Z=a}}]);
//# sourceMappingURL=716.32b48bfb.chunk.js.map