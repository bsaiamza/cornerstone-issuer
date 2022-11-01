"use strict";(self.webpackChunkcornerstone_issuer=self.webpackChunkcornerstone_issuer||[]).push([[869],{5192:function(e,n,t){t.d(n,{H:function(){return r}});t(2791);var i=t(184),r=function(e){var n=e.title;return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)("meta",{charSet:"utf-8"}),(0,i.jsx)("meta",{name:"robots",content:"noindex, follow"}),(0,i.jsx)("meta",{name:"description",content:"Debi Check Query System"}),(0,i.jsx)("meta",{name:"viewport",content:"width=device-width, initial-scale=1, shrink-to-fit=no"}),(0,i.jsxs)("title",{children:["IAMZA | ",n]})]})}},5869:function(e,n,t){t.r(n),t.d(n,{default:function(){return M}});var i,r=t(168),s=t(885),a=t(2791),d=t(5953),l=t(6015),o=t(3118),u=t(5651),c=t(7985),m=t(5803),x=t(4565),h=t(7205),b=t(1508),f=t(6863),j=t(1087),g=t(5192),y=t(1413),p=t(4165),Z=t(5861),_=t(4142),v=t(6650),C=t(1436),w=t(2506),A=t(5206),S=t(1044),k=t(7792),P=t(6571),z=t(9204),I=t(6591),q=t(184),F=I.rd+"/credential",R=I.rd+"/email-credential",N=t(1908),B=function(e){var n={};return N(e.identity_number)||(n.identity_number="Invalid ZA ID number!"),n},D=function(e){var n=e.handleNext,t=(0,_.Z)(),i=(0,a.useState)(!1),r=(0,s.Z)(i,2),l=r[0],o=r[1],u=function(){var e=(0,Z.Z)((0,p.Z)().mark((function e(t){return(0,p.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return o(!0),e.next=3,A.Am.promise(S.ZP.post(F,t).then((function(e){sessionStorage.setItem("credential",e.data.credential),A.Am.success("Credential generated!"),setTimeout((function(){n()}),1e3)})).catch((function(e){A.Am.error(e.response.data.msg)})),{pending:"Generating credential..."});case 3:o(!1);case 4:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}(),c=function(){var e=(0,Z.Z)((0,p.Z)().mark((function e(n){return(0,p.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return o(!0),e.next=3,A.Am.promise(S.ZP.post(R,n).then((function(e){sessionStorage.setItem("credential",e.data.credential),A.Am.success("Emailed credential!")})).catch((function(e){A.Am.error(e.response.data.msg)})),{pending:"Emailing credential..."});case 3:o(!1);case 4:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}();return(0,q.jsxs)(d.ZP,{container:!0,children:[(0,q.jsx)(d.ZP,{item:!0,xs:0,md:2}),(0,q.jsx)(d.ZP,{item:!0,xs:12,md:8,children:(0,q.jsx)(v.Z,{square:!0,elevation:2,sx:{p:3,width:{md:"100%"},backgroundColor:t.palette.mode===I.cp?"#fff":"",borderRadius:5},children:(0,q.jsx)(w.J9,{initialValues:{identity_number:"",names:"",surname:"",gender:"",date_of_birth:"",country_of_birth:"",nationality:"",citizen_status:"",email:""},validate:B,onSubmit:function(e,n){n.resetForm;""===e.email?u(e):c(e)},children:function(e){var n=e.values,t=e.handleChange,i=e.touched,r=e.errors,s=e.setFieldValue;return(0,q.jsxs)(w.l0,{children:[(0,q.jsxs)("div",{children:[(0,q.jsx)(C.Z,{error:i.identity_number&&Boolean(r.identity_number),helperText:i.identity_number&&r.identity_number,id:"identity_number",name:"identity_number",value:n.identity_number,onChange:t,label:"ID Number",sx:{m:"1rem"},required:!0}),(0,q.jsx)(C.Z,{id:"names",name:"names",value:n.names,onChange:t,label:"Names",sx:{m:"1rem"},required:!0})]}),(0,q.jsxs)("div",{children:[(0,q.jsx)(C.Z,{id:"surname",name:"surname",value:n.surname,onChange:t,label:"Surname",sx:{m:"1rem"},required:!0}),(0,q.jsx)(C.Z,{id:"gender",name:"gender",value:n.gender=n.identity_number?n.identity_number.substring(6,7)>4?"Male":"Female":"",onChange:t,label:"Gender",sx:{m:"1rem"},required:!0,disabled:!0})]}),(0,q.jsxs)("div",{children:[(0,q.jsx)(P._,{dateAdapter:k.y,children:(0,q.jsx)(z.M,{label:"Date of Birth",value:n.date_of_birth=n.identity_number?n.identity_number.substring(0,1)>2?"19"+n.identity_number.substring(0,2)+"-"+n.identity_number.substring(2,4)+"-"+n.identity_number.substring(4,6):"20"+n.identity_number.substring(0,2)+"-"+n.identity_number.substring(2,4)+"-"+n.identity_number.substring(4,6):"",onChange:function(e){return s("date_of_birth",e,!0)},renderInput:function(e){return(0,q.jsx)(C.Z,(0,y.Z)((0,y.Z)({id:"date_of_birth",name:"date_of_birth"},e),{},{sx:{m:"1rem",width:"14.5rem"},required:!0}))},disabled:!0})}),(0,q.jsx)(C.Z,{id:"country_of_birth",name:"country_of_birth",value:n.country_of_birth="0"===n.identity_number.substring(10,11)?"RSA":"Other",onChange:t,label:"Country of Birth",sx:{m:"1rem"},required:!0,disabled:!0})]}),(0,q.jsxs)("div",{children:[(0,q.jsx)(C.Z,{id:"nationality",name:"nationality",value:n.nationality="0"===n.identity_number.substring(10,11)?"RSA":"Other",onChange:t,label:"Nationality",sx:{m:"1rem"},required:!0,disabled:!0}),(0,q.jsx)(C.Z,{id:"citizen_status",name:"citizen_status",value:n.citizen_status="0"===n.identity_number.substring(10,11)?"Citizen":"Non-Citizen",onChange:t,label:"Citizen Status",sx:{m:"1rem"},required:!0,disabled:!0})]}),(0,q.jsx)("div",{children:(0,q.jsx)(C.Z,{id:"email",name:"email",type:"email",value:n.email,onChange:t,label:"Email",sx:{m:"1rem"},helperText:(0,q.jsx)(b.Z,{severity:"info",sx:{backgroundColor:"transparent"},children:"Optional: If you want to receive your credential via email."})})}),(0,q.jsx)("div",{children:(0,q.jsx)(h.Z,{variant:"contained",size:"small",type:"submit",sx:{color:"#fff",m:"1rem"},disabled:l,children:"Submit"})})]})}})})}),(0,q.jsx)(d.ZP,{item:!0,xs:0,md:2})]})},E=t(925),O=function(){var e=(0,_.Z)(),n=sessionStorage.getItem("credential");return(0,q.jsx)("div",{style:{backgroundColor:e.palette.mode===I.cp?"":"#F5F5F5",padding:e.palette.mode===I.cp?"":5,borderRadius:e.palette.mode===I.cp?"":5},children:n?(0,q.jsx)(E.Z,{value:n}):"No QR code to scan!"})},Q=[{label:"Enter Details",description:"Please capture your details that will be stored in your credential."},{label:"Scan QR Code",description:"Please scan the QR code with your mobile wallet app to receive your Cornerstone Credential."}],H=(0,f.ZP)(j.OL)(i||(i=(0,r.Z)(["\n\ttext-decoration: none;\n\tcolor: inherit;\n"]))),M=function(){var e=(0,a.useState)(0),n=(0,s.Z)(e,2),t=n[0],i=n[1],r=function(){i((function(e){return e+1}))},f=function(){i((function(e){return e-1}))};return(0,q.jsxs)(q.Fragment,{children:[(0,q.jsx)(g.H,{title:"Credential"}),(0,q.jsxs)(d.ZP,{container:!0,spacing:1,children:[(0,q.jsxs)(d.ZP,{item:!0,xs:12,md:4,sx:{borderRight:{xs:0,md:1},borderBottom:{xs:1,md:0}},children:[(0,q.jsx)(l.Z,{sx:{textAlign:"left",p:1},children:(0,q.jsx)(o.Z,{activeStep:t,orientation:"vertical",children:Q.map((function(e,n){return(0,q.jsxs)(u.Z,{children:[(0,q.jsx)(c.Z,{children:e.label}),(0,q.jsxs)(m.Z,{children:[(0,q.jsx)(x.Z,{children:e.description}),(0,q.jsx)(l.Z,{sx:{mb:2},children:(0,q.jsxs)("div",{children:[n===Q.length-3?(0,q.jsx)(h.Z,{variant:"contained",size:"small",onClick:r,sx:{mt:1,mr:1},children:"Continue"}):"",n===Q.length-1?(0,q.jsx)(h.Z,{variant:"outlined",size:"small",onClick:f,sx:{mt:1,mr:1},children:"Back"}):""]})})]})]},e.label)}))})}),t===Q.length-1&&(0,q.jsx)(q.Fragment,{children:(0,q.jsx)(b.Z,{severity:"info",sx:{m:2,textAlign:"left"},children:(0,q.jsx)(h.Z,{variant:"outlined",size:"small",color:"info",sx:{mt:1,mr:1},children:(0,q.jsx)(H,{to:"/",children:"Return Home"})})})})]}),(0,q.jsxs)(d.ZP,{item:!0,xs:12,md:8,sx:{alignItems:"center",justifyContent:"center",textAlign:"center"},justifyContent:"center",justifyItems:"center",children:[t===Q.length-2&&(0,q.jsx)(D,{handleNext:r}),t===Q.length-1&&(0,q.jsx)(O,{})]})]})]})}}}]);
//# sourceMappingURL=869.23db70be.chunk.js.map