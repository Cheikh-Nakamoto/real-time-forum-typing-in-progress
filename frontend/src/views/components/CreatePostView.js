import { createPost } from "../../api/api.js";
import { getView, setView } from "../../lib/lib.js";
import { Button, Div, Form, TextField, Text, MaterialIcon } from "../elements/index.js";

export class CreatePost {
  constructor() {
    this.path = "/create-post";
    this.id = "createPost";
    setView(this);
    // this.render();
  }
  get element() {
    return new Div({
      id:"create-post-overlay",
      style:{
        top:0,
        position: 'fixed',
        display: "none",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        width: "100vw",
        minHeight: "100vh",
        gap: "1rem",
        zIndex:"1000",
        backgroundColor:"#5c595970",
      },
      children:[
        new Div({
      id: "create-post-container",
      className: "create-post-container",
      style: {
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        width:"55%",
        minWidth: "250px",
        maxWidth:"520px",
        padding:"1rem 2rem",
        gap: "1rem",
        height:"fit-content",
        backgroundColor:"white",
        borderRadius:"10px"


      },
      children: [
        new Div({
          id: "create-post-header",
          className: "create-post-header",
          style: {
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            padding:".5rem",
            width:"100%",
            gap: "1rem",
            fontSize: "1rem",
            fontWeight: "bold",
            justifyContent:"space-between"
          },
          children: [
            new Div({
              children:[
                new Text({
                  text: "Create a post",
                }),
              ]
            }),
            new MaterialIcon({
              iconName: "close",
              style: {
                color: "var(--bs-gray)",
                fontSize: "1.5rem",
                alignSelf:"end",
              },
              listeners:{
              onclick: () => {
                this.hide();
              },}
            }),
          ],
        }),
        new Form({
          id: "create-post-form",
          className: "create-post-form",
          style: {
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            width: "100%",
            height: "100%",
            gap: "1rem",
          },
          children: [
            new TextField({
              id: "create-post-title",
              name:"title",
              className: "create-post-title",
              placeholder: "Title",
              style: {
                width: "100%",
                height: "100%",
                padding: "0.5rem",
                borderRadius: "10px",
                border: "1px solid var(--bs-gray)",
              },
            }),
            new TextField({
              id: "create-post-body",
              name:"body",
              className: "create-post-body",
              placeholder: "Body",
              multiLine:true,
              style: {
                width: "100%",
                minHeight: "300px",
                padding: "1rem",
                borderRadius: "10px",
                border: "1px solid var(--bs-gray)",
              },
            }),
            new Button({
              id: "create-post-submit",
              className: "create-post-submit clickable",
              style: {
                width: "100%",
                height: "100%",
                padding: "0.5rem",
                borderRadius: "10px",
                border: "1px solid var(--bs-gray)",
                backgroundColor: "var(--bs-blue)",
                color: "white",
                fontWeight: "bold",
              },
              children:[
                new Text({
                  text: "Create Post",
                }),
              ],
             
            }),
          ],
          listeners: {
            onsubmit: (event) => {
              event.preventDefault();
              const form = event.target;
              const formData = new FormData(form);
              const data = {};
              formData.forEach((value, key) => {
                data[key] = value;
              });
              createPost(data).then(response=>{
                if (response) {
                  dispatchEvent(new CustomEvent("newPost",{detail:response.post}))
                  this.hide()
                }
              }); 
            },
          },
        }),
      ],
    })
      ]
    }).element;
    
    
  }

  hide(){
    const view = getView("create-post-overlay")
   view.element.style.display = "none";
  }
  show(){
    const view = getView("create-post-overlay")
    view.element.style.display = "flex";
  }
}
