import { Card, Descriptions, Spin } from "antd";
import { LockTwoTone, UnlockTwoTone } from '@ant-design/icons';

function RequsetResult({ userReq, reqLoading }) {
    const { userId, config, duration, encryptRequired, blockNumber, transactionIndex, address } = userReq

    //console.log(userReq)
    return (
        <Spin spinning={reqLoading}>
            <Card title="Your request" extra={encryptRequired ? <LockTwoTone /> : <UnlockTwoTone />}>
                <Descriptions column={2} bordered>
                    <Descriptions.Item label="用户ID" span={2}>{userId}</Descriptions.Item>
                    <Descriptions.Item label="CPU核心数" span={2}>{config}</Descriptions.Item>
                    <Descriptions.Item label="有效期限" span={2}>{duration} days</Descriptions.Item>
                    <Descriptions.Item label="交易地址" span={2}>{address}</Descriptions.Item>
                    <Descriptions.Item label="区块号">{blockNumber}</Descriptions.Item>
                    <Descriptions.Item label="偏移量">{transactionIndex}</Descriptions.Item>
                </Descriptions>
            </Card>
        </Spin>

    )
}

export default RequsetResult;