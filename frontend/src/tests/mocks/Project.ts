import faker from 'faker';

faker.seed(1);

export default function Project(): Project {
  const title = faker.random.word();
  return {
    id: faker.datatype.uuid(),
    title,
    slug: faker.helpers.slugify(title),
    image: faker.image.imageUrl(),
    description: faker.random.words(),
    createdAt: faker.date.soon().toString(),
    updatedAt: faker.date.soon().toString(),
    url: faker.internet.url(),
    git: faker.internet.url(),
    tags: [...new Array(10)].map(() => ProjectTag()),
  };
}

export function ProjectTag() {
  const title = faker.random.word();
  return {
    id: faker.datatype.uuid(),
    title,
    slug: faker.helpers.slugify(title),
    createdAt: faker.date.soon().toString(),
    updatedAt: faker.date.soon().toString(),
  };
}
