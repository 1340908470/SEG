import socket from "@/backend/socket";
import {call} from "@/utils/client";
import auth from "@/backend/auth";

export default function Msg (){
  var s = new WebSocket("ws://localhost:8001/api/test/TestService.TestSocket")
  s.onopen = () => {
    call(auth.UserService.GetProfile, {}).then(
      resp => {
        s.send(resp.User.ID.toString())
      }
    )
  }
  s.onmessage = (e) => {
    alert("sock: " + e.data)
    console.log("sock: " + e.data)
  }


  return <div>11</div>
}
