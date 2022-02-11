import React from 'react';
import Project from '@/tests/mocks/Project';
import Tag from '@/tests/mocks/Tag';
import query from '@/components/Projects/query.graphql';
import { MockedProvider } from '@apollo/client/testing';
import { render } from '@testing-library/react';
import { setMobile } from '@/tests/utils';

import Index from '../pages/index';

const mocks = [
  {
    request: {
      query,
    },
    result: {
      data: {
        tags: [...new Array(10)].map(() => Tag()),
        projects: [...new Array(10)].map(() => Project()),
      },
    },
  },
];

const component = (
  <MockedProvider mocks={mocks}>
    <Index />
  </MockedProvider>
);

it('matches its snapshot', () => {
  const { baseElement } = render(component);
  expect(baseElement).toMatchSnapshot();
});

it('matches its mobile snapshot', () => {
  setMobile();
  const { baseElement } = render(component);
  expect(baseElement).toMatchSnapshot();
});
