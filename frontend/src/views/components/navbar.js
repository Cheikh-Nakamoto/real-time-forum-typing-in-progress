import { logout } from "../../api/api.js";
import { resetOpenedChat } from "../../lib/lib.js";
import { goTo } from "../../lib/pager/navigation.js";
import { Div, Nav, MaterialIcon, TextField, Text, Image } from "../elements/index.js";
import { Logo } from "./logo.js";
export class Navbar {
  constructor() {

    return new Nav({
      className: "navbar",
      children: [
        new Div({
          style:{
            display:"flex",
            flexDirection:'row',
            gap:"2rem",
            justifyContent:"center",
            alignItems:"center"
          },
          children: [
            new Logo(),
            new Div({
              className: "logo-hamburger",
              children: [
                new MaterialIcon({
                  iconName: "menu",
                })
              ]
            }),
          ]
        }),


        new Div({
          className: "search-bar",
          style: {
            width: "35%",
            backgroundColor: "#bed9ec",
            borderRadius: "8px",
            padding: "3px",
          },
          children: [
            new TextField({
              placeholder: "Search here...",
              className: "search-input",
              style: {
                width: "100%",
                border: 'none',
                backgroundColor: "transparent",
                outline: 'none',
                padding: "0 15px",
                color: "var(--bs-gray-dark)"
              }
            }),
            new MaterialIcon({ iconName: "search" }),
          ]
        }),

        new Div({
          className: "icon-bar",
          children: [
            new MaterialIcon({ iconName: "home" }), 
            new Div({
              style:{
                display:"flex",
                flexDirection:'row',
                gap:"0.5rem"
              },
              children:[
                new Image({
                  src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${app.user.username}`,
                  alt: "Author avatar",
                  style: {
                      width: "22px",
                      height: "22px",
                      borderRadius: "50%",
                      backgroundColor: "var(--bs-gray)",
                  },
              }),
                new Div({
                  style:{
                    color:"var(--bs-blue)"
                  },
                  children:[
                    new Text({
                     text: app.user.username
                    })
                  ]
                }),
                new MaterialIcon({
                  className:"logout",
                  iconName:"logout",
                  listeners:{
                    onclick:()=>{
                      logout().then(response=>{
                        if (response) {
                          localStorage.clear();
                          resetOpenedChat()
                          app.user = null
                          goTo("auth")
                        }
                      });
                    }
                  }
                })
              ]
            }),
              ]
            }),
           

      ]
    })
  }
}
