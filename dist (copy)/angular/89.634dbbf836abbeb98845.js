(self.webpackChunkangular=self.webpackChunkangular||[]).push([[89],{1089:(t,s,e)=>{"use strict";e.r(s),e.d(s,{DraftModule:()=>i});var r=e(1116),n=e(7368),u=e(1671),a=e(9425);let c=(()=>{class t{constructor(t,s){this.store=t,this.router=s}ngOnDestroy(){this.authSubs&&this.authSubs.unsubscribe()}ngOnInit(){this.authSubs=this.store.select("auth").subscribe(t=>{t.ID||this.router.navigateByUrl("/admin/login")})}}return t.\u0275fac=function(s){return new(s||t)(n.Y36(u.yh),n.Y36(a.F0))},t.\u0275cmp=n.Xpm({type:t,selectors:[["app-draft"]],decls:2,vars:0,template:function(t,s){1&t&&(n.TgZ(0,"p"),n._uU(1,"draft works!"),n.qZA())},styles:[""]}),t})(),i=(()=>{class t{}return t.\u0275fac=function(s){return new(s||t)},t.\u0275mod=n.oAB({type:t}),t.\u0275inj=n.cJS({imports:[[r.ez,a.Bz.forChild([{path:"",component:c}])]]}),t})()}}]);