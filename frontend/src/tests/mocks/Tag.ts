import faker from 'faker';

export default function Tag() {
  const title = faker.random.word();
  return {
    id: faker.datatype.uuid(),
    title,
    slug: faker.helpers.slugify(title),
    createdAt: faker.date.soon().toString(),
    updatedAt: faker.date.soon().toString(),
    count: faker.datatype.number(),
  };
}
