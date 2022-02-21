type Maybe<T> = T | null;
type InputMaybe<T> = Maybe<T>;
type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

type Asset = {
  __typename?: 'Asset';
  createdAt: Scalars['String'];
  id: Scalars['String'];
  slug: Scalars['String'];
  title: Scalars['String'];
  updatedAt: Scalars['String'];
  url: Scalars['String'];
};

type Project = {
  __typename?: 'Project';
  assets?: Maybe<Array<Maybe<Asset>>>;
  createdAt: Scalars['String'];
  description: Scalars['String'];
  git?: Maybe<Scalars['String']>;
  id: Scalars['String'];
  slug: Scalars['String'];
  tags?: Maybe<Array<Maybe<Tag>>>;
  title: Scalars['String'];
  updatedAt: Scalars['String'];
  url?: Maybe<Scalars['String']>;
};

type Projects = {
  __typename?: 'Projects';
  items?: Maybe<Array<Maybe<Project>>>;
  items_total: Scalars['Int'];
  tags?: Maybe<Array<Maybe<Tag>>>;
  total: Scalars['Int'];
};

type Query = {
  __typename?: 'Query';
  ping?: Maybe<Scalars['String']>;
  projects?: Maybe<Projects>;
  tags?: Maybe<Array<Maybe<Tag>>>;
};


type QueryProjectsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  tags?: InputMaybe<Array<InputMaybe<Scalars['String']>>>;
};

type Tag = {
  __typename?: 'Tag';
  count?: Maybe<Scalars['Int']>;
  createdAt: Scalars['String'];
  id: Scalars['String'];
  slug: Scalars['String'];
  title: Scalars['String'];
  updatedAt: Scalars['String'];
};
