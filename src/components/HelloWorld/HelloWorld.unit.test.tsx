import React from 'react';
import { render } from '@testing-library/react';

import HelloWorld from './';

it('matches its snapshot', () => {
  expect(render(<HelloWorld />).asFragment).toMatchSnapshot();
});
