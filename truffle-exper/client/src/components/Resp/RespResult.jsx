import { Card, Descriptions, Spin } from "antd";


function RespResult({ result, respLoading }) {

    const { vmId, vmName, assignTime, endTime } = result

    return (
        <Spin spinning={respLoading}>
            <Card title="Response Data">
                <Descriptions bordered column={2}>
                    <Descriptions.Item label="虚拟机Id">{vmId}</Descriptions.Item>
                    <Descriptions.Item label="虚拟机名称">{vmName}</Descriptions.Item>
                    <Descriptions.Item label="分配起始时间">{assignTime}</Descriptions.Item>
                    <Descriptions.Item label="分配终止时间">{endTime}</Descriptions.Item>
                </Descriptions>
            </Card>
        </Spin>

    )
}

export default RespResult;