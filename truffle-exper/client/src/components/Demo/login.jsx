import React from 'react';
import { Button, Input, Form } from 'antd';
import { publicKey, privateKey } from '../../utils/encryption';
import useEth from '../../contexts/EthContext/useEth';
const Login = ({setUserId, setPubKey}) => {
    const { state } = useEth()
    const onFinish = (values) => {
        setUserId(state.accounts)
        setPubKey(publicKey)
        console.log('Success:', values);
        sessionStorage["userId"] = state.accounts
        sessionStorage["pubKey"] = publicKey
        sessionStorage["priKey"] = privateKey
    };

    return (
        <div style={{
            paddingTop: `2rem`
        }}>
            <Form
                name="login_form"
                layout="inline"
                onFinish={onFinish}
                initialValues={{
                    userId: 'user1'
                }}
                size="large"
            >
                {/* <Form.Item
                    name="userId"
                    label="UserId"
                >
                    <Input
                        defaultValue="user1"
                        placeholder='Please input your userId'
                    />
                </Form.Item> */}
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        获取列表信息
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
};
export default Login;

