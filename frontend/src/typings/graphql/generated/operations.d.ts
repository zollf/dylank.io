type WorkQueryVariables = Exact<{
  tags?: InputMaybe<Array<InputMaybe<Scalars['String']>> | InputMaybe<Scalars['String']>>;
  offset?: InputMaybe<Scalars['Int']>;
  limit?: InputMaybe<Scalars['Int']>;
}>;


type WorkQuery = { projects?: { total: number, itemsTotal: number, items: Array<{ id: string, title: string, slug: string, updatedAt: string, insertedAt: string, shortDescription?: string | null, previewLink?: string | null, gitLink?: string | null, pageContent?: string | null, tags?: Array<{ id?: string | null, slug?: string | null, title?: string | null } | null> | null } | null>, tags: Array<{ id?: string | null, slug?: string | null, title?: string | null, count?: number | null } | null> } | null };

type TagFragment = { id?: string | null, slug?: string | null, title?: string | null, insertedAt: string, updatedAt: string };
