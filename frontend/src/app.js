
import { Div } from "./views/elements/index.js";
import { ViewPager } from "./lib/pager/view-pager.js";
import { ContentPage } from "./views/pages/ContentPage.js";
import { AuthPage } from "./views/pages/index.js";
import {  handleWSEvent, setWSConnection } from "./api/api.js";

export class App {
  constructor(options) {
    this.container = options.container
    this.currentPath = options.currentPath
    this.wsConnection
    this.user = options.user

    history.pushState("app", "", this.currentPath)
  }
  setContentView() {
    return new ViewPager({
      id: "pager1",
      defaultPage: /\/auth\/\w+/.test(this.currentPath) ? { id: "auth", page: AuthPage } : { page: ContentPage, id: "contentPage" },
      pages: [{ id: "auth", page: AuthPage }, { page: ContentPage, id: "contentPage" }],
      container: new Div(),
      locPathname: location.pathname
    })
  }

  run() {

    // addEventListener("load", (event) => {
    const content = this.setContentView()
    this.container.appendChild(content.container.element)
    // })

    addEventListener("logged", (event) => {
      setWSConnection()
      
      if (!this.user) {
        const user = event.detail.user
        this.user = user
      }


      this.wsConnection.addEventListener("open", (event) => {
      })

      this.wsConnection.addEventListener("message", (event) => {
        handleWSEvent(event.data)
      })
      this.wsConnection.onclose = function (event) {
        // if (event.wasClean) {
        //   alert(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
        // } else {
        //   // e.g. server process killed or network down
        //   // event.code is usually 1006 in this case
        //   alert('[close] Connection died');
        // }
      };

    })




  }
}
