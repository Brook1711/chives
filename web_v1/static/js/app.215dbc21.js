/*! chives project */(function(e){function t(t){for(var r,a,c=t[0],i=t[1],u=t[2],d=0,f=[];d<c.length;d++)a=c[d],Object.prototype.hasOwnProperty.call(o,a)&&o[a]&&f.push(o[a][0]),o[a]=0;for(r in i)Object.prototype.hasOwnProperty.call(i,r)&&(e[r]=i[r]);p&&p(t);while(f.length)f.shift()();return l.push.apply(l,u||[]),n()}function n(){for(var e,t=0;t<l.length;t++){for(var n=l[t],r=!0,c=1;c<n.length;c++){var i=n[c];0!==o[i]&&(r=!1)}r&&(l.splice(t--,1),e=a(a.s=n[0]))}return e}var r={},o={app:0},l=[];function a(t){if(r[t])return r[t].exports;var n=r[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,a),n.l=!0,n.exports}a.m=e,a.c=r,a.d=function(e,t,n){a.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},a.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},a.t=function(e,t){if(1&t&&(e=a(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(a.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)a.d(n,r,function(t){return e[t]}.bind(null,r));return n},a.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return a.d(t,"a",t),t},a.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},a.p="static/";var c=window["webpackJsonp"]=window["webpackJsonp"]||[],i=c.push.bind(c);c.push=t,c=c.slice();for(var u=0;u<c.length;u++)t(c[u]);var p=i;l.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("7a23"),o=Object(r["h"])("data-v-6bf07146");Object(r["g"])("data-v-6bf07146");var l=Object(r["d"])("div",{id:"dplayer",class:"play-root"},null,-1);Object(r["f"])();var a=o((function(e,t,n,o,a,c){return Object(r["e"])(),Object(r["c"])(r["a"],null,[l,Object(r["d"])("input",{type:"file",id:"input",accept:"video/*",onChange:t[1]||(t[1]=function(){return c.playSelectedFile&&c.playSelectedFile.apply(c,arguments)})},null,32)],64)})),c=(n("2b3d"),n("d3b7"),n("3ca3"),n("ddb0"),n("f7a5")),i=n.n(c),u={name:"App",data:function(){return{msg:"hello vue!  ",dplayer:"",fileURL:""}},mounted:function(){console.log("dp_app")},methods:{playSelectedFile:function(){var e=document.getElementById("input"),t=e.files[0],n=URL.createObjectURL(t);this.fileURL=n,console.log("file selected"),console.log("dp_app");var r=new i.a({container:document.getElementById("dplayer"),video:{url:n},danmaku:{api:"http://161.35.234.230:1207/",addition:["http://161.35.234.230:1207/v3/bilibili?aid=80266688&cid=137358410"]},screenshot:!0});console.log(r),this.dplayer=r},a:function(){console.log("clicked")}}};u.render=a,u.__scopeId="data-v-6bf07146";var p=u,d=Object(r["b"])(p).mount("#app");console.log(d)}});
//# sourceMappingURL=app.215dbc21.js.map