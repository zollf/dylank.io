import faker from 'faker';

faker.seed(1);

export default function Project(): Project {
  const title = faker.random.word();
  return {
    id: faker.datatype.uuid(),
    title,
    slug: faker.helpers.slugify(title),
    shortDescription: faker.random.words(),
    insertedAt: faker.date.soon().toString(),
    updatedAt: faker.date.soon().toString(),
    previewLink: faker.internet.url(),
    gitLink: faker.internet.url(),
    tags: [...new Array(10)].map(() => ProjectTag()),
  };
}

export function ProjectTag() {
  const title = faker.random.word();
  return {
    id: faker.datatype.uuid(),
    title,
    slug: faker.helpers.slugify(title),
    insertedAt: faker.date.soon().toString(),
    updatedAt: faker.date.soon().toString(),
  };
}
