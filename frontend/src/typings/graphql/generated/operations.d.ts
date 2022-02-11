type WorkQueryVariables = Exact<{
  tags?: InputMaybe<Array<InputMaybe<Scalars['String']>> | InputMaybe<Scalars['String']>>;
  offset?: InputMaybe<Scalars['Int']>;
  limit?: InputMaybe<Scalars['Int']>;
}>;


type WorkQuery = { projects?: Maybe<(
    Pick<Projects, 'total' | 'items_total'>
    & { items?: Maybe<Array<Maybe<(
      Pick<Project, 'id' | 'title' | 'slug' | 'image' | 'description' | 'createdAt' | 'updatedAt' | 'url' | 'git'>
      & { tags?: Maybe<Array<Maybe<Pick<Tag, 'id' | 'slug' | 'title' | 'createdAt' | 'updatedAt'>>>> }
    )>>>, tags?: Maybe<Array<Maybe<Pick<Tag, 'id' | 'slug' | 'title' | 'createdAt' | 'updatedAt' | 'count'>>>> }
  )> };

type TagFragment = Pick<Tag, 'id' | 'slug' | 'title' | 'createdAt' | 'updatedAt'>;
