import { getUsersStatus } from "../../../api/api.js";
import { getView } from "../../../lib/lib.js";
import { Div } from "../../elements/index.js";
import { ListView } from "../ListView.js";
import { StatusItemView } from "./status-item-view.js";

export class StatusPanel {
  constructor() {
    this.usersStatus = new ListView({
      id: "userList",
      itemView: StatusItemView,
      provider: getUsersStatus,
    });

    addEventListener("newUser", (event) => {
      this.usersStatus.refresh();
    });

    addEventListener("newMessage", (event) => {
      const data = event.detail;
      let user;
      let statusItmView;
      if (data.From != app.user.username) {
        user = getView(`user-status${data.From}`).user;
        statusItmView = getView(`status-item-${data.From}`);
      } else {
        user = user = getView(`user-status${data.To}`).user;
        statusItmView = getView(`status-item-${data.To}`);
      }
      this.usersStatus.removeItemView(statusItmView);
      this.usersStatus.prependItem(user);
    });

    return new Div({
      className: "status-panel",
      style: {
        flex: 1,
        display: "flex",
        flexDirection: "column",
        gap: "1.5rem",
        padding: "20px 10px",
        height: "100%",
        backgroundColor: "var(--bs-white)",
        boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
      },
      children: [
        new Div({
          style: {
            width: "fit-content",
            height: "100%",
            position: "fixed",
            display: "flex",
            flexDirection: "column",
            gap: "1rem",
            overflowY: "scroll",
          },
          children: [
            new Div({
              className: "chats",
              children: [this.usersStatus.listContainer],
            }),
          ],
        }),
      ],
    });
  }
}
