webpackJsonp([1],{"2z6H":function(t,e){},"3BF7":function(t,e){},"7zck":function(t,e){},NHnr:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=a("7+uW"),n={render:function(){var t=this.$createElement,e=this._self._c||t;return e("v-app",{staticStyle:{height:"100%"}},[e("v-toolbar",{attrs:{id:"toolbar",app:""}},[e("v-toolbar-title",[this._v("Chat Application")])],1),this._v(" "),e("v-content",{staticClass:"pt-0 fill-height"},[e("router-view")],1)],1)},staticRenderFns:[]};var r=a("VU/8")({name:"App"},n,!1,function(t){a("3BF7")},null,null).exports,o=a("/ocq"),i={data:function(){return this.$store.commit("setName",""),{name:""}},methods:{login:function(){this.$store.commit("setName",this.name),console.log(this.$store.state.name),this.$router.push("/chat")}}},c={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-container",{attrs:{fluid:"","fill-height":""}},[a("v-layout",{attrs:{"align-center":"","justify-center":""}},[a("v-flex",{attrs:{xs12:"",sm8:"",md4:""}},[a("v-card",{staticClass:"elevation-12"},[a("v-toolbar",{attrs:{dark:"",color:"blue-grey lighten-2"}},[a("v-toolbar-title",[t._v("Login")])],1),t._v(" "),a("v-card-text",[a("v-form",[a("v-text-field",{attrs:{"prepend-icon":"person",name:"login",label:"Your Nickname",type:"text"},model:{value:t.name,callback:function(e){t.name=e},expression:"name"}})],1)],1),t._v(" "),a("v-card-actions",[a("v-spacer"),t._v(" "),a("v-btn",{attrs:{color:"primary"},on:{click:function(e){t.login()}}},[t._v("Login")])],1)],1)],1)],1)],1)},staticRenderFns:[]},l=a("VU/8")(i,c,!1,null,null,null).exports,u=a("BO1k"),v=a.n(u),m={name:"Top",created:function(){try{var t=new WebSocket("/ws");t.onclose=function(){this.$router.push("/")},t.onmessage=function(t){var e=JSON.parse(t.data);this.messages.push(e),console.log(e)},this.ws=t}catch(t){alert(t),this.$router.push("/")}},methods:{send:function(){var t=this.message;this.message="",this.ws.send({user:this.$store.state.name,message:t})}},computed:{getMessages:function(){var t=0,e=!0,a=!1,s=void 0;try{for(var n,r=v()(this.messages);!(e=(n=r.next()).done);e=!0){n.value.id=t,t++}}catch(t){a=!0,s=t}finally{try{!e&&r.return&&r.return()}finally{if(a)throw s}}return this.messages}},data:function(){return{messages:[],message:"",ws:null}}},h={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"fill-height"},[a("div",{attrs:{id:"chat-history"}},t._l(t.getMessages,function(e){return a("div",{key:e.id,staticClass:"chat-element"},[a("v-layout",{attrs:{column:""}},[a("div",{staticClass:"chat-text body-1"},[t._v(t._s(e.user))]),t._v(" "),a("div",{staticClass:"chat-text subheading"},[t._v(t._s(e.message))])])],1)})),t._v(" "),a("div",{attrs:{id:"chat-bar"}},[a("v-layout",{attrs:{row:"","align-center":""}},[a("v-flex",{attrs:{xs10:""}},[a("v-text-field",{attrs:{label:"Message","auto-grow":"",autofocos:"",rows:"2"},model:{value:t.message,callback:function(e){t.message=e},expression:"message"}})],1),t._v(" "),a("v-flex",{attrs:{xs2:""}},[a("v-btn",{attrs:{color:"primary"},on:{click:function(e){t.send()}}},[t._v("Send")])],1)],1)],1)])},staticRenderFns:[]};var f=a("VU/8")(m,h,!1,function(t){a("2z6H")},"data-v-60acce89",null).exports,d=a("NYxO");s.default.use(d.a);var p=new d.a.Store({state:{name:""},mutations:{setName:function(t,e){t.name=e}}});s.default.use(o.a);var g=new o.a({routes:[{path:"/",name:"Top",component:l},{path:"/chat",name:"Chat",component:f}]});g.beforeEach(function(t,e,a){"/"!==t.path?(null==p.name&&a("/"),a()):a()});var _=g,b=a("3EgV"),x=a.n(b);a("7zck"),a("csSS"),a("gJtD");s.default.use(x.a),s.default.config.productionTip=!1,new s.default({el:"#app",router:_,store:p,components:{App:r},template:"<App/>"})},csSS:function(t,e){},gJtD:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.49a977f69314efc1f89f.js.map