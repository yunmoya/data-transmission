import { useEffect, useState } from "react";
import useEth from "../../contexts/EthContext/useEth";
import NoticeNoArtifact from "./NoticeNoArtifact";
import NoticeWrongNetwork from "./NoticeWrongNetwork";
import RespResult from "./RespResult";
import axios from "axios"
import NoticeNoUser from "./NoticeNoUser";
import { decrypt } from "../../utils/encryption";
import Web3 from "web3";

function Resp({ resp, respLoading, setRespLoading }) {
  const { state } = useEth();
  const [result, setResult] = useState({
    userId: "",
    vmId: "",
    vmName: "",
    assignTime: "",
    endTime: ""
  })

  const [errorMsg, setErrorMsg] = useState("")

  const [userId, setUserId] = useState(sessionStorage["userId"]);

  const respForm =
    <div>
      <h2>响应信息</h2>
      <br />
      <RespResult result={result} respLoading={respLoading} />
    </div>;

  useEffect(() => {
    console.log("The resp is: ", resp)
    console.log("now is ", new Date().getTime(), "verify resp")
    if (resp.data !== '') {
      axios.post('http://localhost:8080/verify', {
        msg: resp.data,
        hash: resp.hash,
        signature: resp.signature
      })
        .then(function (response) {
          console.log(response);
          if (response.data.IsValid) {
            var msgBytes = Web3.utils.hexToBytes(resp.data)
            var msg = Buffer.from(msgBytes)
            if(resp.encryptRequired){
              console.log(Buffer.from(msg))
              console.log("now is ", new Date().getTime(), "decrypt resp")
              msg = decrypt(sessionStorage["priKey"], Buffer.from(msg))
              console.log("Decrypted Msg:", msg.toString())
            }
            var obj = JSON.parse(msg.toString())
            console.log(obj)
            setResult({
              vmId: obj.vmId,
              vmName: obj.vmName,
              assignTime: obj.assignTime,
              endTime: obj.endTime
            })
            setErrorMsg("")
          } else {
            setErrorMsg(response.data.Desc)
          }
        })
        .catch(function (error) {
          setErrorMsg(error)
          console.log(error)
        });
        setRespLoading(false)
    }

  }, [resp, setResult])

  return (
    <div>
      {
        !state.artifact ? <NoticeNoArtifact /> :
          !state.contract ? <NoticeWrongNetwork /> :
            userId === "" || userId === undefined ? <NoticeNoUser /> :
              errorMsg !== "" ? <span>{errorMsg}</span> :
                respForm
      }
    </div>
  );
}

export default Resp;
