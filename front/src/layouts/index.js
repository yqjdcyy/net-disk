import React from 'react';
import { Menu, Icon } from 'antd';
const { SubMenu } = Menu;

// todo: menu.jump
class NetDiskMenu extends React.Component {
  handleClick = e => {
    window.location.href = e.key;
  };

  render() {
    return (
      <Menu onClick={this.handleClick} theme="dark" mode="horizontal">
        <SubMenu
          key="files"
          title={
            <span>
              <Icon type="file-protect" />
              Files
            </span>
          }
        >
          <Menu.ItemGroup title="TMP">
            <Menu.Item key="/file?path=/data/tmp/img">img</Menu.Item>
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

export default NetDiskMenu;
