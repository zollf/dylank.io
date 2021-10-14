import React from 'react';
import { render } from '@testing-library/react';

import GoodByeWorld from './';

it('matches its snapshot', () => {
  expect(render(<GoodByeWorld />).asFragment).toMatchSnapshot();
});
