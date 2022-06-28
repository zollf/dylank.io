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

type Project = {
  __typename?: 'Project';
  gitLink?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  insertedAt: Scalars['String'];
  pageContent?: Maybe<Scalars['String']>;
  previewLink?: Maybe<Scalars['String']>;
  shortDescription?: Maybe<Scalars['String']>;
  slug: Scalars['String'];
  tags?: Maybe<Array<Maybe<Tag>>>;
  title: Scalars['String'];
  updatedAt: Scalars['String'];
};

type ProjectInterface = {
  __typename?: 'ProjectInterface';
  items: Array<Maybe<Project>>;
  itemsTotal: Scalars['Int'];
  tags: Array<Maybe<TagInterface>>;
  total: Scalars['Int'];
};

type RootQueryType = {
  __typename?: 'RootQueryType';
  projects?: Maybe<ProjectInterface>;
  tags?: Maybe<Array<Maybe<Tag>>>;
};


type RootQueryTypeProjectsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  tags?: InputMaybe<Array<InputMaybe<Scalars['String']>>>;
};

type Tag = {
  __typename?: 'Tag';
  id?: Maybe<Scalars['ID']>;
  insertedAt: Scalars['String'];
  slug?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  updatedAt: Scalars['String'];
};

type TagInterface = {
  __typename?: 'TagInterface';
  count?: Maybe<Scalars['Int']>;
  id?: Maybe<Scalars['ID']>;
  insertedAt: Scalars['String'];
  slug?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  updatedAt: Scalars['String'];
};
