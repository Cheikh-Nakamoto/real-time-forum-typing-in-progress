import { Div } from "../../elements/index.js";
import { MenuItem } from "./menu-item.js";
import { getView } from "../../../lib/lib.js";

export class Menu {
  constructor() {
    return new Div({
      className: "menu",
      constraints: {
        top: function () {
          const elem = getView("header");
          return elem.element.clientHeight + "px";
        },
      },
      style: {
        display: "flex",
        flexDirection: "column",
        alignItems: "baseline",
        justifyContent: "space-between",
        gap: "1.5rem",
        width: "fit-content",
        padding: "20px 0",
        position: "fixed",
        alignSelf: "flex-start",

      },
      children: [
        new MenuItem({
          title: "Home",
          iconName: "home",
        }),
        new MenuItem({
          title: "Explore",
          iconName: "explore",
        }),
        new MenuItem({
          title: "About",
          iconName: "info",
        }),
      ],
    });
  }
}
