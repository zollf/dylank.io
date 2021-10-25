import React from 'react';
import { render } from '@testing-library/react';
import { setMobile } from '@/tests/utils';

import Index from './';

it('matches its snapshot', () => {
  const { baseElement } = render(<Index />);
  expect(baseElement).toMatchSnapshot();
});

it('matches its mobile snapshot', () => {
  setMobile();
  const { baseElement } = render(<Index />);
  expect(baseElement).toMatchSnapshot();
});
