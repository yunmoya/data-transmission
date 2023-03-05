import { useEffect, useState } from "react";
import { Row, Col } from "antd"
import useEth from "../../contexts/EthContext/useEth";
import Title from "./Title";
import RequestForm from "./RequestForm";
import NoticeNoArtifact from "./NoticeNoArtifact";
import NoticeWrongNetwork from "./NoticeWrongNetwork";
import Login from "./login";
import RequsetResult from "./RequestResult";

function Demo({ userReq, reqLoading, setReqLoading, setRespLoading }) {
  const { state } = useEth();
  const [value, setValue] = useState("?");
  const [userId, setUserId] = useState(sessionStorage["userId"]);
  const [pubKey, setPubKey] = useState(sessionStorage["pubKey"]);

  const demo =
    <div>
      <br />
      <Row>
        <Col span={12}>
          {/* <ContractBtns setValue={setValue} /> */}
          <RequestForm setValue={setValue} userId={userId} pubKey={pubKey} setReqLoading={setReqLoading} setRespLoading={setRespLoading} />
        </Col>
        <Col span={12}>
          <RequsetResult userReq={userReq} reqLoading={reqLoading} />
        </Col>

      </Row>
    </div>;

  return (
    <div>
      <Title />
      {
        !state.artifact ? <NoticeNoArtifact /> :
          !state.contract ? <NoticeWrongNetwork /> :
            userId === "" || userId === undefined ? <Login setUserId={setUserId} setPubKey={setPubKey} /> :
              demo
      }
    </div>
  );
}

export default Demo;
