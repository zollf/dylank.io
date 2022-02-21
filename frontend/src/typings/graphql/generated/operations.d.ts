type WorkQueryVariables = Exact<{
  tags?: InputMaybe<Array<InputMaybe<Scalars['String']>> | InputMaybe<Scalars['String']>>;
  offset?: InputMaybe<Scalars['Int']>;
  limit?: InputMaybe<Scalars['Int']>;
}>;


type WorkQuery = { projects?: { total: number, items_total: number, items?: Array<{ id: string, title: string, slug: string, description: string, createdAt: string, updatedAt: string, url?: string | null, git?: string | null, assets?: Array<{ id: string, slug: string, title: string, createdAt: string, updatedAt: string, url: string } | null> | null, tags?: Array<{ id: string, slug: string, title: string, createdAt: string, updatedAt: string } | null> | null } | null> | null, tags?: Array<{ id: string, slug: string, title: string, createdAt: string, updatedAt: string, count?: number | null } | null> | null } | null };

type TagFragment = { id: string, slug: string, title: string, createdAt: string, updatedAt: string };
