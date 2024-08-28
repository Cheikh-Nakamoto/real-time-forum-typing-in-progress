import { getPosts } from "../../api/api.js";
import { debounce, getView, throttle } from "../../lib/lib.js";
import { CreatePost, ListView, PostCard } from "../components/index.js";
import { Div, Image, Text, TextField } from "../elements/index.js";

export class MainPage {
  constructor(options) {
    this.id = "mainPage"
    this.pagerId = options.pagerId
    this.title = "Main Page"
    this.path = "/"


  }
  get element() {
    const postList = new ListView({
      id: "postList",
      itemView: PostCard,
      provider: getPosts,
    });

    addEventListener("newPost", (event) => {
      const post = event.detail;
      postList.prependItem(post);
    })

    addEventListener('scroll',() => {
      const scrollTop = document.scrollingElement.scrollTop;

      const isAtBottom = throttle(()=>{
       return scrollTop >= (document.scrollingElement.scrollHeight * 0.80)
      },500)
        if (isAtBottom()) {
          postList.fetchMoreItems()
        }
    });

    return new Div({
      className: "main-page",
      style: {
        display: "flex",
        flexDirection: "column",
        paddingTop: "1rem",
        height: "100%",
        gap: "1rem",
      },
      children: [
        new Div({
          className: "new-post-card",
          style: {
            backgroundColor: "var(--bs-white)",
            width: "100%",
            display: "flex",
            flexDirection: "column",
            alignItems: "start",
            padding: "1rem",
            borderRadius: "10px",
          },
          children: [
            new Div({
              children: [
                new Text({
                  text: "Create Post",
                }),
              ],
            }),
            new Div({
              style: {
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                padding: "15px",
                width: "100%",
                height: "100%",
                gap: "1rem",
              },
              children: [
                new Image({
                  src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${app.user.username}`,
                  alt: "New Post",
                  style: {
                    width: "48px",
                    height: "48px",
                    borderRadius: "50%",
                    backgroundColor: "var(--bs-gray)",
                  },
                }),
                new TextField({
                  placeholder: "What's on your mind?",
                  style: {
                    width: "100%",
                    height: "3.5rem",
                  },
                  listeners: {
                    onfocus: () => {
                      const createPostView = getView("createPost")
                      createPostView.show();
                    }
                  }
                }),
              ],
            }),
          ],
        }),
        postList.listContainer
      ],
    }).element
  }
}
