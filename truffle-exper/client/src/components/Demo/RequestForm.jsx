import { useState } from "react";
import { Form, InputNumber, Select, Checkbox, Button, Input } from "antd";
import useEth from "../../contexts/EthContext/useEth";
import axios from "axios";
import { decrypt } from "../../utils/encryption";
import Web3 from "web3";

function RequestForm({ setValue, userId, pubKey, setReqLoading, setRespLoading }) {
    const { state: { contract, accounts } } = useEth();

    const { Option } = Select

    const [form] = Form.useForm();

    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 10 },
    };

    const defaultValues = {
        userId: userId,
        config: '8',
        duration: '10',
        encryptRequired: false
    };

    const onFinish = async (values) => {
        console.log(values)
        setReqLoading(true)
        setRespLoading(true)
        const { config, duration } = values
        var encryptRequired = true
        console.log("now is ", new Date().getTime(), "I send a userreq tx")
        encryptRequired ?
            await contract.methods.requestVMWithEncryption(userId, config, duration, pubKey).send({ from: accounts[0] }) :
            await contract.methods.requestVMWithoutEncryption(userId, config, duration).send({ from: accounts[0] });
    }

    const testOnClick = async (values) => {
        axios.get('http://localhost:8080/test', {
            params: {
              key: pubKey
            }
          })
          .then(function (response) {
            console.log(response)
            var msgBytes = Web3.utils.hexToBytes("0x"+response.data)
            console.log(msgBytes)
            var data = decrypt(sessionStorage["priKey"], Buffer.from(msgBytes))
            console.log(data.toString())
          })
          .catch(function (error) {
            console.log(error);
          })  
    }

    return (
        <>
            <Form {...layout} form={form} size='large' initialValues={defaultValues} name="requestForm" onFinish={onFinish}>
            <Form.Item name="userId" label="用户ID" rules={[{ required: true }]}>
                <Input disabled/>
            </Form.Item>
            <Form.Item name="config" label="CPU核心数" rules={[{ required: true }]}>
                <Select
                    placeholder="选择CPU核心数"
                >
                    <Option value="8">8 cores</Option>
                    <Option value="16">16 cores</Option>
                    <Option value="32">32 cores</Option>
                    <Option value="64">64 cores</Option>
                    <Option value="128">128 cores</Option>
                    <Option value="256">256 cores</Option>
                </Select>
            </Form.Item>

            <Form.Item name="duration" label="有效期限" rules={[{ required: true }]}>
                <InputNumber controls addonAfter="day" />
            </Form.Item>

            

            <Button type="primary" htmlType="submit">
                提交
            </Button>

            {/* <Button type="primary" onClick={testOnClick}>
                Test
            </Button> */}

        </Form>

        </>
        
    );
}

export default RequestForm;
