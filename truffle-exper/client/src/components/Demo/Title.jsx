import useEth from "../../contexts/EthContext/useEth";

function Title() {
  const { state } = useEth();
  return <h2>{state.accounts}，你好</h2>;
}

export default Title;
