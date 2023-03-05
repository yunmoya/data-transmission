import { EthProvider, RespEthProvider } from "./contexts/EthContext";
import { useState } from "react";
import Intro from "./components/Intro/";
import Setup from "./components/Setup";
import Demo from "./components/Demo";
import Resp from "./components/Resp"
import Footer from "./components/Footer";
import "./App.css";

function App() {

  const [userReq, setUserReq] = useState({
    userId: "",
    config: 0,
    duration: 0,
    encryptRequired: false,
    blockNumber: 0,
    transactionIndex: 0,
    address: ""
  })

  const [resp, setResp] = useState({
    userId: "",
    hash: "",
    signature: "",
    data: "",
    encryptRequired: false
  })

  const [reqLoading, setReqLoading] = useState(false)

  const [respLoading, setRespLoading] = useState(false)

  return (
    <div id="App" >
      <div className="container">
        <Intro />
        <hr />
        {/* <Setup />
        <hr /> */}
        <EthProvider setUserReq={setUserReq} setReqLoading={setReqLoading}>
          <Demo userReq={userReq} reqLoading={reqLoading} setReqLoading={setReqLoading} setRespLoading={setRespLoading}/>
        </EthProvider>
        <hr />
        <RespEthProvider setResp={setResp}>
          <Resp resp={resp} respLoading={respLoading} setRespLoading={setRespLoading} />
        </RespEthProvider>
        <hr />
        <Footer />
      </div>
    </div>
  );
}

export default App;
