import { EventType } from "../../../api/api.js";
import { getView, setView, throttle } from "../../../lib/lib.js";
import { Div, Image, Text } from "../../elements/index.js";

export class StatusItemView {
  constructor(user) {
    this.user = user;
    this.id = `user-status${user.username}`;
    this.unreadMsg = user.unread_count || 0;

    addEventListener("unreadMsg", (event) => {
      const u = event.detail;
      const count = u.unread_count;
      const badge = getView(`unread-badge-${this.id}`);
      if (this.user.username === u.username) {
        if (count > 0) {
          badge.element.innerText = count > 9 ? "+9" : count;
          badge.show();
        } else {
          badge.hide();
        }
      }
    });

    setView(this);
    return new Div({
      id: `status-item-${this.user.username}`,
      className: "status-item",
      style: {
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        width: "fit-content",
        height: "fit-content",
        gap: "1rem",
      },
      children: [
        new Div({
          style: {
            position: "relative",
          },
          children: [
            new Div({
              id: `unread-badge-${this.id}`,
              className: "unread-badge",
              style: {
                display: this.unreadMsg > 0 ? "flex" : "none",
                position: "absolute",
                top: "5px",
                right: "3px",
                width: "15px",
                height: "15px",
                backgroundColor: "red",
                borderRadius: "50%",
                color: "white",
                fontWeight: "bold",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                textAlign: "center",
                fontSize: "10px",
                padding: "2px",
              },
              children: [
                new Text({
                  text: this.unreadMsg > 9 ? "+9" : this.unreadMsg,
                }),
              ],
            }),
            new Image({
              src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${this.user.username}`,
              alt: "Author avatar",
              style: {
                width: "48px",
                height: "48px",
                borderRadius: "50%",
                backgroundColor: "var(--bs-gray)",
              },
            }),
            new Div({
              className: "dot",
              id: `${this.user.username}status-dot`,
              style: {
                position: "absolute",
                bottom: "5px",
                left: "3px",
                width: "10px",
                height: "10px",
                borderRadius: "50%",
                backgroundColor: user.status === "online" ? "green" : "gray",
              },
            }),
          ],
        }),
        new Div({
          style: {
            display: "flex",
            flexDirection: "column",
            alignItems: "start",
            justifyContent: "center",
          },
          children: [
            new Div({
              style: {
                color: "var(--bs-blue)",
                fontSize: "14px",
              },
              children: [
                new Text({
                  text: this.user.username,
                }),
              ],
            }),
            new Div({
              id: `${this.user.username}-status-text`,
              style: {
                className: "user-status-text",
                color: "var(--bs-blue)",
                fontSize: "12px",
              },
              children: [
                new Text({
                  text: this.user.status,
                }),
              ],
            }),
          ],
        }),
      ],
      listeners: {
        onclick: throttle(() => {
          const newEvent = new CustomEvent("chatOpened", { detail: this.user });
          dispatchEvent(newEvent);

          this.user.unread_count = 0;
          dispatchEvent(new CustomEvent("unreadMsg", { detail: this.user }));
          const wsEvent = {
            type: EventType.WS_READ_EVENT,
            time: new Date(Date.now()).toString(),
            username: this.user.username,
          };
          app.wsConnection.send(JSON.stringify(wsEvent));
        }, 500),
      },
    });
  }
}
