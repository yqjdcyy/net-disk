import React from 'react';
import { Menu, Icon } from 'antd';
const { SubMenu } = Menu;

export class NetDiskMenu extends React.Component {
  state = {
    current: 'files',
  };

  handleClick = e => {
    console.log('click ', e);
    this.setState({
      current: e.key,
    });
  };

  render() {
    return (
      <Menu onClick={this.handleClick} selectedKeys={[this.state.current]} mode="horizontal">
        <SubMenu key="files"
        title={
            <span>
                <Icon type="file-protect" />
                Files
            </span>
        }>
            <Menu.ItemGroup title="CDN">
                <Menu.Item key="cdn:img">img</Menu.Item>
            </Menu.ItemGroup>
            <Menu.ItemGroup title="Tmp">
                <Menu.Item key="tmp:img">img</Menu.Item>
            </Menu.ItemGroup>
        </SubMenu>
        <Menu.Item key="setting">
            <Icon type="setting" />
            Settings
        </Menu.Item>
      </Menu>
    );
  }
}