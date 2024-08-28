import { EventType, checkSession, getMessages } from "../../api/api.js";
import { getView, remOpenedChat, remView, throttle } from "../../lib/lib.js";
import {
  Div,
  Image,
  MaterialIcon,
  Text,
  TextField,
} from "../elements/index.js";
import { ListView } from "./ListView.js";
import { MessageView } from "./MessageView.js";
import { TypingIndicator } from "./typing-indicator/TypingIndicator.js";

export class ChatView {
  constructor(prop) {
    this.id ="chat" + prop.user.username
    this.chat = prop.chat;
    this.recipient = prop.user;
    this.typingIndicator = new TypingIndicator()

    this.messageList = new ListView({
      id: `messageList${this.chat.chat_id}`,
      itemView: MessageView,
      provider: getMessages,
      providerParams: {
        chatId: this.chat.chat_id,
      },
      style: {
        flexDirection: "column-reverse",
        overflowY: "scroll",
        height: "100%",
        width: "100%",
        overflowX: "hidden",
      },
    });

    addEventListener("newMessage", (event) => {
      const message = event.detail.Data;
      this.messageList.prependItem(message);
    });
    let typingstatus;
    addEventListener("typing", (event) => {
      const data = event.detail;
      const typingNotification = getView(
        `${data.From}-typingstatus`
      );
      // const messContainer = getView(`list-${this.messageList.id}-container`)
      clearTimeout(typingstatus);
      if (typingNotification) {
        // const typingIndicator = new TypingIndicator()
        // messContainer.addChild(typingIndicator)
        this.showTypingIndicator()
        typingNotification.element.style.display = "flex";
        typingstatus = setTimeout(() => {
          console.log("timeout");
          typingNotification.element.style.display = "none";
          // messContainer.removeChild(typingIndicator)
          this.hideTypingIndicator()

        }, 2000);

      } else {  

      }
    });
  }

  get element() {
    return new Div({
      id: this.id,
      className: "chat-container",
      style: {
        display: "flex",
        flexDirection: "column",
        position: "relative",
        position: "relative",
        width: "360px",
        boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
        maxHeight: "55vh",
        maxHeight: "55vh",
        backgroundColor: "aliceblue",
      },
      children: [
        new Div({
          id: `chatHeader${this.recipient.username}`,
          className: "chat-header",
          style: {
            backgroundColor: "rgb(190, 217, 236)",
            display: "flex",
            flexDirection: "row",
            width: "100%",
            padding: "0.2rem .5rem",
            justifyContent: "space-between",
            alignItems: "center",
            borderTopLeftRadius: "10px",
            borderTopRightRadius: "10px",
            transition: "max-height 0.5s ease-out",
          },
          children: [
            new Div({
              style: {
                display: "flex",
                flexDirection: "row",
                width: "100%",
                display: "flex",
                flexDirection: "row",
                width: "100%",
                position: "relative",
                alignItems: "center",
                gap: ".5rem",
              },
              children: [
                new Div({
                  className: "img-dot",
                  style: {
                    position: "relative",
                  },
                  children: [
                    new Image({
                      src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${this.recipient.username}`,
                      alt: "Author avatar",
                      style: {
                        width: "32px",
                        height: "32px",
                        borderRadius: "50%",
                        backgroundColor: "var(--bs-gray)",
                      },
                    }),
                    new Div({
                      className: "dot",
                      id: `chat-${this.recipient.username}-status-dot`,
                      style: {
                        bottom: "5px",
                        left: "3px",
                        width: "10px",
                        height: "10px",
                        borderRadius: "50%",
                        backgroundColor:
                          this.recipient.status === "online" ? "green" : "gray",
                        position: "absolute",
                      },
                    }),
                  ],
                }),
                new Div({
                  className: "chat-infos",
                  children: [
                    new Text({ text: this.recipient.username }),
                    new Div({
                      id: `${this.recipient.username}-typingstatus`,
                      className: "typed",
                      style: {
                        display: "none",
                        color: "var(--bs-blue)",
                        fontSize: "1rem",
                        bottom: "0",
                      },
                      children: [new Text({ text: "is typing..." })],
                    }),
                  ],
                }),
              ],

              listeners: {
                onclick: () => {
                  this.toggleDisplay();
                },
              },
            }),
            new MaterialIcon({
              iconName: "close",
              className: "chat-close",
              listeners: {
                onclick: () => {
                  const view = getView("chat" + this.recipient.username);
                  const elem = view.element;
                  elem.parentNode.removeChild(elem);
                  remOpenedChat(view.id);
                  const chatContainer = getView("chatsContainer");
                  chatContainer.chatViews.splice(
                    chatContainer.chatViews.indexOf(elem),
                    1
                  );
                },
              },
            }),
          ],
        }),
        new Div({
          id: `chatContainer${this.chat.chat_id}`,
          className: "messages-container",
          style: {
            display: "flex",
            width: "100%",
            padding: "0.2rem 0.5rem",
            height: "55vh",
            maxHeight: "55vh",
            transition: "max-height 0.5s ease-out",
            flexDirection: "column",
            gap: "1.5rem",
            paddingBottom: "1.1rem",
          },
          children: [
            new Div({
              style: {
                display: "flex",
                flexDirection: "column",
                width: "100%",
                height: "75%",
                justifyContent: "bottom",
                // marginBottom: '"3rem',
              },
              children: [this.messageList.listContainer,this.typingIndicator.element],
            }),
            new Div({
              id: "msgtyperWrapper",
              style: {
                // flex:"1",
                width: "100%",
                bottom: "10px",
                padding: "0 1rem",
                alignSelf: "end",
                display: "flex",
                flexDirection: "row",
                gap: "1rem",
                justifyContent: "center",
                alignItems: "center",
              },
              children: [
                new TextField({
                  id: "msg-input",
                  placeholder: "type your message",
                  multiLine: true,
                  style: {
                    fontFamily: "Open Sans",
                    height: "34px",
                    width: "100%",
                    outline: "none",
                    borderRadius: "15px",
                    border: "1px solid var(--bs-blue)",
                    padding: "10px",
                  },
                  listeners: {
                    oninput: (e) => this.handleInput(e),
                  },
                }),
                new MaterialIcon({
                  iconName: "send",
                  style: {
                    color: "var(--bs-white)",
                    backgroundColor: "var(--bs-blue)",
                    borderRadius: "10px",
                    padding: "5px",
                  },
                  listeners: {
                    onclick: () => {
                      this.send();
                    },
                    onkeydown: (e) => {
                      if ((e.key = "Enter")) {
                        this.send();
                      }
                    },
                  },
                }),
              ],
            }),
          ],
        }),
      ],
    }).element;
  }

  get getInput() {
    return getView("msg-input").element.value;
  }

  async send() {
    const text = this.getInput.trim();

    if (text != "") {
      checkSession().then((response) => {
        if (response) {
          const wsEvent = {
            type: EventType.WS_MESSAGE_EVENT,
            to: this.recipient.username,
            content: text,
            time: new Date(Date.now()).toString(),
            chatId: this.chat.chat_id,
          };
          app.wsConnection.send(JSON.stringify(wsEvent));
          this.resetInput();
        }
      });
    }
  }

  resetInput() {
    getView("msg-input").element.value = "";
  }

  toggleDisplay() {
    let div = getView(`chatContainer${this.chat.chat_id}`).element;
    if (div.style.maxHeight === "0px") {
      div.style.maxHeight = "55vh";
      div.style.height = "55vh";
      setTimeout(() => {
        div.style.display = "flex";
      }, 300);
    } else {
      div.style.maxHeight = "0px";
      div.style.height = "0px";
      div.style.display = "none";
      // div.style.visibility = "hidden"
    }
  }
  handleInput = throttle(() => {
    console.log("typing");
    const wsEvent = this.generateWsEvent();
    this.sendMessage(wsEvent);
  }, 300);

  generateWsEvent() {
    return {
      type: EventType.WS_TYPING_EVENT,
      to: this.recipient.username,
      from: app.user.username,
      content: "",
      time: new Date(Date.now()).toString(),
    };
  }
  sendMessage(wsEvent) {
    app.wsConnection.send(JSON.stringify(wsEvent));
  }
  // showTypingnotification() {

  //   if (typingNotification !== undefined) {
  //     const elem = typingNotification.element;
  //     console.log(elem);
  //     elem.style.display === "none"
  //       ? (elem.style.display = "flex")
  //       : (elem.style.display = "none");
  //   }
  // }

  showTypingIndicator(){
    const elem = document.getElementById(this.typingIndicator.element.id)
    elem.style.display= "flex"
  }
  hideTypingIndicator(){
    const elem = document.getElementById(this.typingIndicator.element.id)
    elem.style.display= "none"
  }
}
