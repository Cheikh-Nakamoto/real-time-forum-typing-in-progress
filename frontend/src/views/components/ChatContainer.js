import { getChatByUser } from "../../api/api.js";
import { getOpenedChat, getView, remOpenedChat, setOpenedChat, setView } from "../../lib/lib.js";
import { Div } from "../elements/index.js";
import { ChatView } from "./chatView.js";

export class ChatContainer {
  constructor() {
    this.id = "chatsContainer";
    this.chatViews = [];

    addEventListener("chatOpened", (event) => {
      const user = event.detail;

      if (!getOpenedChat("chat" + user.username)) {
        if (this.chatViews.length == 1) {
          const view = this.chatViews.shift();
          // document.getElementById(view.id)?.remove();
          remOpenedChat(view.id);
        }
        
        getChatByUser(user.username).then((chat) => {

          if (document.getElementById("chat" + user.username)) {
            return;
          }
          const chatView = new ChatView({ chat, user });
          setOpenedChat(chatView);
          const chatViewElement = chatView.element;
          const container = getView("chats-container").element;
          container.appendChild(chatViewElement);
          this.chatViews.push(chatViewElement);
        });
      }
    });
    setView(this);
  }

  get element() {
    return new Div({
      className: "chats-container",
      id: "chats-container",
      style: {
        display: "flex",
        flexDirection: "row-reverse",
        alignItems: "end",
        position: "fixed",
        bottom: 0,
        right: "8%",
        minWidth: "fit-content",
        maxWidth: "45%",
        maxHeight: "55%",
        gap: "1rem",
      },
      children: [],
    }).element;
  }
}
