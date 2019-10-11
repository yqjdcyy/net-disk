import React from 'react';
import { Icon, Card } from 'antd';

class FileCard extends React.Component {
  handleClick = e => {
    console.log('');
  };

  render() {
    return (
      <Card
        hoverable
        style={{ width: 240 }}
        cover={
          <img alt="example" src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" />
        }
        actions={[<Icon type="fullscreen" key="fullscreen" />, <Icon type="delete" key="delete" />]}
      ></Card>
    );
  }
}

export default FileCard;
