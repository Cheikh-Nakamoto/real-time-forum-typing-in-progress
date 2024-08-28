import { formatDate } from "../../lib/lib.js";
import { Div, Text } from "../elements/index.js";

export class MessageView {
    constructor(message) {
        this.message = message;
    }

    get element() {
        return new Div({
            className: `message-container ${this.message.isSender ? 'outgoing':'incoming'}`,
            style: {
                display: "flex",
                gap:"0.5rem",
                maxWidth: "85%",
                overflowWrap: "anywhere"
            },
            children: [
                new Div({
                    style: {
                        color: "var(--bs-gray)",
                        fontWeight: "bold",
                        alignSelf: "end",
                        fontSize:"12px"
                        
                    },
                    children: [
                        new Text({
                            text: formatDate(new Date(this.message.created_at)) ,

                        })
                    ]
                }),
                new Div({
                    className:"msg-content",
                    style: {
                        fontWeight: "bold",
                        color: "var(--bs-white)",
                        padding: "1rem 1.2rem",
                        fontFamily:"Open Sans",

                    },
                    children: [
                        new Text({
                            text: this.message.content,

                        })
                    ]
                }),
               
            ]
        }).element
    }
}