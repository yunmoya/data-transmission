import React, { useReducer, useCallback, useEffect } from "react";
import Web3 from "web3";
import EthContext from "./EthContext";
import { reducer, actions, initialState } from "./state";

function EthProvider({ children, setUserReq, setReqLoading }) {
  const [state, dispatch] = useReducer(reducer, initialState);

  const init = useCallback(
    async artifact => {
      if (artifact) {
        const web3 = new Web3(Web3.givenProvider || "ws://localhost:7545");
        // const web3 = new Web3(Web3.givenProvider || "ws://goerli.infura.io/v3/547bcf201d264561b95c50860d2dddff");
        const accounts = await web3.eth.requestAccounts();
        const networkID = await web3.eth.net.getId();
        const { abi } = artifact;
        let address, contract;
        try {
          address = artifact.networks[networkID].address;
          contract = new web3.eth.Contract(abi, address);
        } catch (err) {
          console.error(err);
        }
        dispatch({
          type: actions.init,
          data: { artifact, web3, accounts, networkID, contract }
        });

        contract.events.NewRequest([], (error, result) => {
          if (!error) {
            // console.log(result)
            const data = result.returnValues
            if (data.userId === sessionStorage["userId"]) {
              setUserReq({
                userId: data.userId,
                config: data.config,
                duration: data.duration,
                encryptRequired: data.encryptRequired,
                blockNumber: result.blockNumber,
                transactionIndex: result.transactionIndex,
                address: result.address
              })
              setReqLoading(false)
            }
          } else {
            console.log(error)
          }
        })
      }
    }, [setUserReq]);

  useEffect(() => {
    const tryInit = async () => {
      try {
        const artifact = require("../../contracts/UserRequestContract.json");
        init(artifact);
      } catch (err) {
        console.error(err);
      }
    };

    tryInit();
  }, [init]);

  useEffect(() => {
    const events = ["chainChanged", "accountsChanged"];
    const handleChange = () => {
      init(state.artifact);
    };

    events.forEach(e => window.ethereum.on(e, handleChange));
    return () => {
      events.forEach(e => window.ethereum.removeListener(e, handleChange));
    };
  }, [init, state.artifact]);

  return (
    <EthContext.Provider value={{
      state,
      dispatch
    }}>
      {children}
    </EthContext.Provider>
  );
}

export default EthProvider;
