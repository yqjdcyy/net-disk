import React from 'react';
import { Menu, Icon } from 'antd';
const { SubMenu } = Menu;

class NetDiskMenu extends React.Component {

  handleClick = e => {
    console.log(e.key)
    console.log(e.url)
  };

  render() {
    return (
      <Menu 
        onClick={this.handleClick} 
        theme="dark"
        mode="horizontal">
          
        <SubMenu key="files"
        title={
            <span>
                <Icon type="file-protect" />
                Files
            </span>
        }>
          <Menu.ItemGroup title="TMP">
              <Menu.Item key="tmp:img" url="/file?path=/data/tmp/img">img</Menu.Item>
          </Menu.ItemGroup>
            <Menu.ItemGroup title="CDN">
                <Menu.Item key="cdn:img" url="/file?path=/data/cdn/img">img</Menu.Item>
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
