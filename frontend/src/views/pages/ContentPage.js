import { DefaultLayout } from "../layout/default_layout.js";
import { Header, Div } from "../elements/index.js";
import { Navbar, Menu, StatusPanel, SuggestionPanel, ChatContainer,CreatePost } from "../components/index.js";
import { MainPage } from "./main-page.js";
import { setPage } from "../../lib/lib.js";
export class ContentPage {
  constructor(options) {
    this.id = "contentPage"
    this.pagerId = options.pagerId
    this.title = "Home"
    this.path = "/"
    this.CreatePostView = new CreatePost()

    setPage(this)

  }
  get element() {

    return new DefaultLayout({
      style: {
        width: '100%',
        height: '100%',
        position:'relative'
      },
      children: [
        new Header({
          id: "header",
          className: "layout-header",
          style: {
            position: "sticky",
            top: 0,
            zIndex: 1000,
          },
          children: [
            new Navbar()
          ]
        }),

        new Div({
          className: "layout-main",
          style: {
            width: '100%',
            height: '100%',
            display: "flex",
            flexDirection: "row",
            gap: "1.5rem"
          },
          children: [
            new Div({
              className: "leftPanel",
              style: {
                flex: 1.5,
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                backgroundColor: 'var(--bs-white)',
                boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
                // top:`${getView("header").clientHeight}`
              },
              children: [
                new Menu()
              ],
            }),
            new Div({
              className: "mainPanel",
              style: {
                flex: 6,
                height: "100%",

              },
              children: [
                new MainPage({ pagerId: "pager2" }),

              ]
            }),
            new Div({
              className: "rightPanel",
              style: {
                position: "relative",
                flex: 4,
                display: "flex",
                flexDirection: "row",
                boxShadow: "-20px 0 15px -23px rgba(0,0,0,0.1)",
                boxShadow: "-20px 0 15px -23px rgba(0,0,0,0.1)",
                gap: "1rem",
              },
              children: [
                new SuggestionPanel(),
                new StatusPanel(),

              ]
            })
          ],
        }),
        new ChatContainer(),
        this.CreatePostView
      ],
    }).element
  }
}