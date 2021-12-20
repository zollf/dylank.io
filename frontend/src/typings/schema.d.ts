interface Tag {
  id: string;
  slug: string;
  title: string;
  createdAt: string;
  updatedAt: string;
  count: number;
}

interface Project {
  title: string;
  slug: string;
  image: string;
  description: string;
  createdAt: string;
  updatedAt: string;
  url: string;
  git: string;
  tags: Array<{
    id: string;
    slug: string;
    title: string;
    createdAt: string;
    updatedAt: string;
  }>;
}
