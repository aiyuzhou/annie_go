"use strict";

import React from "react";
import createReactClass from "create-react-class";
import {Form, Input, Button, message} from "antd";

const FormItem = Form.Item;

let Register = createReactClass({
  render: function () {
    return (
      <section className="register">
        <div className="annie_logo"></div>
        <div className="register_form">
          <form>
            <p className="form-item">
              <label>
                NickName:
                <input type="text" name="name"/>
              </label>
            </p>
            <p className="form-item">
              <label>
                Email:
                <input type="text" name="name"/>
              </label>
            </p>
            <p className="form-item">
              <label>
                Phone Number:
                <input type="text" name="name"/>
              </label>
            </p>
            <p className="form-item">
              <label>
                Password:
                <input type="text" name="name"/>
              </label>
            </p>
            <p className="form-item">
              <label>
                Confirm Password:
                <input type="text" name="name"/>
              </label>
            </p>
            <p className="form-item">
              <Button type="submit">Register</Button>
            </p>
          </form>
        </div>

        <form action="/v1/users/register" method="post">

        手机号码：<input type="text" name="phone"/><br />
        姓名：<input type="text" name="name"/><br />
        密码：<input type="text" name="password"/><br />

        <input type="submit" value="提交"/>
        <input type="reset"/>

        </form>
      </section>
    )
  }
});

export default Register;
