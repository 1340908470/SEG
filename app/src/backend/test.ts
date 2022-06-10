import { Endpoint } from '.';

interface TestReq {
  name: string
}

interface TestRes {
  message:string
}

export default {
  TestService: {
    Test: 'test/TestService.Test' as Endpoint<
      TestReq,
      TestRes
    >,
  },
};
