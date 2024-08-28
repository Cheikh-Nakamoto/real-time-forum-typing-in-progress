import { setView } from "../../lib/lib.js";
import { Div, Image,Text } from "../elements/index.js";

export class CommentView {
  constructor(comment) {
    this.id = comment.comment_id;
    setView(this);
    return new Div({
      id: `comment-container-${this.id}`,
      className: "comment-container",
      style: {
        width: "100%",
        display: "flex",
        flexDirection: "column",
        alignItems: "start",
        padding: "1rem",
        borderRadius: "10px",
        boxShadow: "0px 10px 15px -3px rgba(0,0,0,0.1)",
      },
      children: [
        new Div({
          id: `comment-info-${this.id}`,
          className: "comment-header",
          style: {
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            padding: "5px",
            width: "100%",
            height: "100%",
            gap: "0.5rem",
          },
          children: [
            new Div({
              style: {
                color: "var(--bs-blue)",
                fontSize:"small"
              },
              children: [new Text({ text: comment.username})],
            }),
          ],
        }),
        new Div({
            style:{
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                width: "100%",
  
            },
          children: [
            new Div({
                children:[
                    new Image({
                        src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${comment.username}`,
                        alt: "Author avatar",
                        style: {
                          width: "32px",
                          height: "32px",
                          borderRadius: "50%",
                          backgroundColor: "var(--bs-gray)",
                        },
                    })
                ]
            }),
            new Div({
              id: `comment-body-${this.id}`,
              className: "comment-body",
              style: {
                overflowWrap: "anywhere",
                padding: "5px 15px",
                width: "100%",
                gap: "1rem",
                textAlign: "left",
                marginBottom: "0.5rem",
              },
              children: [new Text({ text: comment.body})],
            }),
          ],
        }),
      ],
    });
  }
}
