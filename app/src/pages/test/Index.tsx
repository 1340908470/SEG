import {Button, Input, Form} from "antd";
import {useState} from "react";
import {call} from "@/utils/client";
import test from "@/backend/test";

export default function Index() {
  return (
    <div>
      <h1>Hello world</h1>
      <Form
        name="basic"
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
        initialValues={{ remember: true }}
        onFinish={(values)=>{
          call(test.TestService.Test, {
            name: values.name
          }).then(resp => {
            alert(resp.message)
          })
        }}
        autoComplete="off"
      >
        <Form.Item
          name="name"
        >
          <Input />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit">
            发送
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}
