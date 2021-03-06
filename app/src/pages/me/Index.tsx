import auth from '@/backend/auth';
import { call } from '@/utils/client';
import Avatar from '@/components/Avatar';
import { Badge, Button, Col, Drawer, Row, Space } from 'antd';
import { useAsync, useSetState } from 'react-use';
import Item from './components/Item';
import { Female, Identification, Male } from './components/Svg';

export default function Me() {
  const [state, setState] = useSetState({ quiting: false });

  const { value: profile } = useAsync(async () => {
    const { User, Student } = await call(auth.UserService.GetProfile, {});
    return { ...User, ...Student };
  });

  return (
    <>
      <Item
        route="/me/account"
        label={
          <Row
            align="middle"
            style={{
              padding: '2em 0',
              fontSize: '16px',
              borderTop: '1px solid #f5f5f5',
              borderBottom: '1px solid #f5f5f5',
            }}
            wrap={false}
          >
            <Col>
              <Avatar
                src={profile?.AvatarURI}
                id={profile?.ID.toString()}
                head={profile?.NickName}
                size={64}
              />
            </Col>
            <Col>
              <span style={{ fontSize: '1.5em', padding: '0.1em' }}>
                {profile?.NickName}
              </span>
              {(() => {
                if (profile?.Gender === true) {
                  return (
                    <Male
                      style={{
                        color: '#52C1D5',
                      }}
                    />
                  );
                } else if (profile?.Gender === false) {
                  return (
                    <Female
                      style={{
                        color: '#EEA5B6',
                      }}
                    />
                  );
                }
              })()}
            </Col>
          </Row>
        }
      ></Item>
      <Space
        direction="vertical"
        size="middle"
        style={{ width: '100%', marginTop: '2em' }}
      >
        <div>
          <Item label="??????">{profile?.RealName}</Item>
          {profile ? (
            <>
              <Item label="??????">{profile.University}</Item>
              <Item label="??????">{profile.School}</Item>
              <Item label="??????">
                {profile.TrustedID || profile.UntrustedID}
              </Item>
            </>
          ) : (
            <>
              <Item label="????????????" route="/me/identify">
                <Identification />
              </Item>
            </>
          )}
          <Item label="????????????">
            {profile?.PhoneNumber.toString().slice(2)}
          </Item>
        </div>
        <Item label="????????????" route="/me/events">
          <span
            style={{
              backgroundColor: 'red',
              color: 'white',
              borderRadius: '1em',
              padding: '0.02em 0.5em 0.2em 0.5em',
              fontSize: '0.8em',
            }}
          >
            new
          </span>
        </Item>
        <Item label="??????" route="/me/setting">
          <Badge dot offset={[-5, 3]}></Badge>
        </Item>
      </Space>
      <div
        style={{
          padding: '2em',
        }}
      >
        <Button
          block
          size="large"
          type="primary"
          onClick={() => {
            setState({ quiting: true });
          }}
        >
          ????????????
        </Button>
      </div>
      <Drawer
        bodyStyle={{
          textAlign: 'center',
          padding: '0',
        }}
        closable={false}
        drawerStyle={{
          height: 'auto',
        }}
        onClose={() => {
          setState({ quiting: false });
        }}
        placement="bottom"
        title={
          <div
            style={{
              textAlign: 'center',
            }}
          >
            ??????????????????????????????
          </div>
        }
        footer={
          <div
            style={{
              color: '#0079FE',
              textAlign: 'center',
            }}
            onClick={() => {
              setState({ quiting: false });
            }}
          >
            ??????
          </div>
        }
        visible={state.quiting}
      >
        <div
          style={{ color: 'red', padding: '1em' }}
          onClick={() => {
            setState({ quiting: false });
          }}
        >
          ????????????
        </div>
      </Drawer>
    </>
  );
}
