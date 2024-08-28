import { getPostComments, postReact } from "../../api/api.js";
import { getView } from "../../lib/lib.js";
import {
  Button,
  Div,
  Image,
  MaterialIcon,
  Text,
  TextField,
} from "../elements/index.js";
import { CommentView } from "./CommentView.js";
import { Divider } from "./divider.js";
import { ListView } from "./ListView.js";
import { postComment } from "../../api/api.js";

export class PostCard {
  constructor(postObject) {
    const commentList = new ListView({
      id: `comment-list-${postObject.post_id}`,
      itemView: CommentView,
      provider: getPostComments,
      providerParams: {
        postid: postObject.post_id,
      },
      style: {
        height: "300px",
        overflowY: "scroll",
      },
    });

    return new Div({
      id: postObject.post_id,
      className: "post-card",
      style: {
        backgroundColor: "var(--bs-white)",
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
          className: "post-header",
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
            new Image({
              src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${postObject.username}`,
              alt: "Author avatar",
              style: {
                width: "48px",
                height: "48px",
                borderRadius: "50%",
                backgroundColor: "var(--bs-gray)",
              },
            }),
            new Div({
              className: "post-info",
              style: {
                display: "flex",
                flexDirection: "column",
                alignItems: "start",
                padding: "15px",
                width: "100%",
                height: "100%",
                // gap: "1rem",
              },
              children: [
                new Div({
                  children: [new Text({ text: postObject.username })],
                }),
                new Div({
                  style: {
                    fontFamily: "Open Sans",
                    fontWeight: "light",
                    fontSize: "12px",
                  },
                  children: [new Text({ text: postObject.Age })],
                }),
              ],
            }),
          ],
        }),
        new Div({
          className: "post-title",
          style: {
            fontSize: "large",
            fontWeight: "bold",
            padding: "5px 15px",
            width: "100%",
            gap: "1rem",
            textAlign: "left",
            marginBottom: "0.5rem",
          },
          children: [new Text({ text: postObject.title })],
        }),
        new Div({
          className: "post-content",
          style: {
            overflowWrap: "anywhere",
            marginBottom: "25px",
          },
          children: [new Text({ text: postObject.body })],
        }),
        new Div({
          style: {
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            width: "100%",
            gap: "1rem",
            justifyContent: "center",
          },
          children: [
            new Divider({
              width: "90%",
            }),
            new Div({
              className: "post-footer",
              style: {
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                width: "100%",
                height: "100%",
                gap: "0.5rem",
                justifyContent: "space-between",
              },
              children: [
                new Div({
                  className: "post-likes",
                  style: {
                    display: "flex",
                    flexDirection: "row",
                    alignItems: "center",
                    padding: "5px",
                    width: "100%",
                    height: "100%",
                    gap: "0.2rem",
                  },
                  children: [
                    new Div({
                      id: `voteCount-${postObject.post_id}`,
                      style: {
                        color: "var(--bs-gray)",
                      },
                      children: [new Text({ text: postObject.Votes })],
                    }),
                    new Div({
                      style: {
                        padding: "1px",
                        color: "var(--bs-gray)",
                      },
                      children: [new Text({ text: "likes" })],
                    }),
                    new Div({
                      style: {
                        display: "flex",
                        flexDirection: "row",
                        alignItems: "center",
                        padding: "5px",
                        gap: "0.5rem",
                      },
                      children: [
                        new MaterialIcon({
                          id: `likeIcon-${postObject.post_id}`,
                          iconName: "thumb_up",
                          className: `${
                            postObject.UserReact === "LIKE"
                              ? "like react"
                              : "react"
                          }`,
                          listeners: {
                            onclick: () => {
                              postReact(postObject.post_id, "LIKE").then(
                                (response) => {
                                  if (response) {
                                    document
                                      .getElementById(
                                        `likeIcon-${postObject.post_id}`,
                                      )
                                      .classList.toggle("like");

                                    const dislikeButton =
                                      document.getElementById(
                                        `dislikeIcon-${postObject.post_id}`,
                                      );

                                    if (
                                      dislikeButton.classList.contains(
                                        "dislike",
                                      )
                                    ) {
                                      dislikeButton.classList.remove("dislike");
                                    }

                                    document.getElementById(
                                      `voteCount-${postObject.post_id}`,
                                    ).innerText = response.votes;
                                  }
                                },
                              );
                            },
                          },
                        }),
                        new MaterialIcon({
                          iconName: "thumb_down",
                          id: `dislikeIcon-${postObject.post_id}`,
                          className: `${
                            postObject.UserReact === "DISLIKE"
                              ? "dislike react"
                              : "react"
                          }`,
                          listeners: {
                            onclick: () => {
                              postReact(postObject.post_id, "DISLIKE").then(
                                (response) => {
                                  if (response) {
                                    document
                                      .getElementById(
                                        `dislikeIcon-${postObject.post_id}`,
                                      )
                                      .classList.toggle("dislike");

                                    const likeButton = document.getElementById(
                                      `likeIcon-${postObject.post_id}`,
                                    );

                                    if (likeButton.classList.contains("like")) {
                                      likeButton.classList.remove("like");
                                    }
                                    document.getElementById(
                                      `voteCount-${postObject.post_id}`,
                                    ).innerText = response.votes;
                                  }
                                },
                              );
                            },
                          },
                        }),
                      ],
                    }),
                  ],
                }),
                new Div({
                  className: "post-comments",
                  style: {
                    display: "flex",
                    flexDirection: "row",
                    alignItems: "center",
                    padding: "5px",
                    height: "100%",
                    gap: "0.2rem",
                    justifyContent: "center",
                  },
                  children: [
                    new MaterialIcon({
                      iconName: "comment",
                      className: "react",
                    }),
                    new Div({
                      children: [new Text({ text: postObject.CommentsCount })],
                      style: {
                        padding: "1px",
                        color: "var(--bs-gray)",
                      },
                    }),
                    new Div({
                      children: [new Text({ text: "comments" })],
                      style: {
                        padding: "1px",
                        color: "var(--bs-gray)",
                      },
                    }),
                  ],
                  listeners: {
                    onclick: () => {
                      const elem = document.getElementById(
                        `post-comments-container-${postObject.post_id}`,
                        );
                      elem.style.display =
                        elem.style.display == "flex" ? "none" : "flex";

                    },
                  },
                }),
              ],
            }),
          ],
        }),
        new Div({
          id: `post-comments-container-${postObject.post_id}`,
          className: "post-comments-container",
          style: {
            display: "none",
            flexDirection: "column",
            width: "100%",
            padding: "5%",
            gap: "1rem",
            justifyContent: "center",
          },
          children: [
            new Div({
              id: `comment-input-wrapper-${postObject.post_id}`,
              className: "comment-input-wrapper",
              style: {
                display: "flex",
                flexDirection: "column",
              },
              children: [
                new TextField({
                  id: `comment-input-${postObject.post_id}`,
                  className: "comment-input",
                  placeholder: "Write a comment",
                  style: {
                    width: "100%",
                    height: "55px",
                    padding: "5px",
                    borderRadius: "5px",
                    boxShadow: "0px 10px 15px -3px rgba(0,0,0,0.1)",
                    outline:"none" ,
                    border:"none",
                    backgroundColor:"aliceblue"
                  },
                }),
                new Button({
                  id: `comment-submit-${postObject.post_id}`,
                  className: "comment-submit",
                  style: {
                    margin: "5px",
                    alignSelf: "end",
                    padding: "5px",
                    borderRadius: "10px",
                    boxShadow: "0px 10px 15px -3px rgba(0,0,0,0.1)",
                    backgroundColor: "var(--bs-blue)",
                    color: "var(--bs-white)",
                  },
                  children: [new Text({ text: "Submit" })],
                  listeners: {
                    onclick: () => {
                      const text = document.getElementById(
                          `comment-input-${postObject.post_id}`,
                        ).value
                      postComment(
                        postObject.post_id,
                        {body:text}
                      ).then(async(response)=>{
                        if(response){
                          commentList.page = 0
                          await commentList.fetch()
                        }
                      })
                    },
                  },
                }),
              ],
            }),
            commentList.listContainer,
          ],
        }),
      ],
    });
  }

  get input(){
    return getView(`comment-input-${postObject.post_id}`).element.value
  }
}
