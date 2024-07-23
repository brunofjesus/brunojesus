---
title: "Learning Symfony"
date: "2021-12-04"
summary: "My approach to learning Symfony"
description: "My approach to learning Symfony"
toc: true
readTime: true
autonumber: true
math: true
tags: ["programming","php"]
showTags: false
---

## 👋 Introduction

In a not so distant past I was occasionaly coding in PHP, as the time passed I got more into Java and Spring, leaving PHP behind.

I'm trying to refresh my knownledge and gather new skills in the process, so I decided to take an Udemy course on Symfony. As I am taking the course I will be leaving some notes here that you can use for you as well.

As always, if you find something wrong there feel free to suggest changes.

## 🏗️ Start a project

### Skeleton

#### Using the Symfony CLI

**Create a new `--full` project if builing a traditional web application:**
```bash
symfony new my_project_name --full
```

**Omit the `--full` when creating a microservice, console app or an API**
```bash
symfony new my_project_name
```

#### Using composer

**Use the `website-skeleton` if creating a traditional web application:**

```bash
composer create-project symfony/website-skeleton my_project_name
```

**Use the `skeleton` if creating a microservice:**
```bash
composer create-project symfony/skeleton my_project_name
```

### Extra project dependencies

**Install doctrine ORM:**

```bash
composer require symfony/orm-pack
```

You might want to install the `doctrine/annotations` package:
```bash
composer require doctrine/annotations
```

**Install a serializer:**

```bash
composer require serializer
```

## 💻 Commands

**List routes:**

```bash
php bin/console debug:router
```

**Create an entity:**

```bash
php bin/console make:entity
```

**Create an user entity:**

```bash
php bin/console make:user
```

**Create a migration:**

```bash
php bin/console make:migration
```

> 📌 Will include changes you made on entities automatically.

**Run migrations:**

```bash
php bin/console doctrine:migrations:migrate
```

## 💾 Fetching data

**Using the repository:**

```php
/**
 * @Route("/post/{id}", name="blog_by_id", requirements={"id"="\d+"})
 */
public function post(int $id): JsonResponse
{
    $repository = $this->getDoctrine()->getRepository(BlogPost::class);

    return $this->json(
        $repository->find($id)
    );
}
```

**Using @ParamConverter implicitly:**

```php
/**
 * @Route("/post/{id}", name="blog_by_id", requirements={"id"="\d+"})
 */
public function post(BlogPost $post): JsonResponse
{
    //Automatically gets the repository and does find($id);
    return $this->json($post);
}
```

> 📌 Notice that in the signature we replaced **int $id** with **BlogPost $post**
Symfony uses the **{id}** in the **@Route** to know which parameter to filter by, being the equivalent of `findOneBy(['id' => <contents of {id}>]);`

**Using `@ParamConverter` explicitly:**

```php
/**
 * @Route("/post/{id}", name="blog_by_id", requirements={"id"="\d+"})
 * @ParamConverter("post", class="App:BlogPost")
 */
public function post($post): JsonResponse
{
    //Automatically gets the repository and does find($id);
    return $this->json($post);
}
```

> 📌 In this example we removed the type **BlogPost** from the signature and did the mapping using an annotation.

**Using `@ParamConverter` with manual mapping:**

```php
/**
 * @Route("/post/{**slug**}", name="blog_by_slug")
 * @ParamConverter("post", options={"mapping": {"slug": "slug"}})
 */
public function postBySlug(BlogPost $post): JsonResponse
{
    //Automatically gets the repository and does findOneBy(['slug' => $slug]);
    return $this->json($post);
}
```

> 📌 In the mapping, the first parameter should match the **route**, while the second one should match the **entity**.

## 🌱 Seeding fake data

**Install the `doctrine-fixtures-bundle` package:**
```bash
composer require --dev doctrine/doctrine-fixtures-bundle
```

**Create the fixture:**
```php
<?php

namespace App\DataFixtures;

use App\Entity\BlogPost;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;

class AppFixtures extends Fixture
{
    private UserPasswordHasherInterface $userPasswordHasher;

    public function __construct(UserPasswordHasherInterface $userPasswordHasher)
    {
        $this->userPasswordHasher = $userPasswordHasher;
    }

    public function load(ObjectManager $manager): void
    {
        $this->loadUsers($manager);
        $this->loadBlogPosts($manager);
    }

    public function loadBlogPosts(ObjectManager $manager)
    {
        /** @var User $user */
        $user = $this->getReference('user_admin');

        $blogPost = new BlogPost();
        $blogPost->setTitle('A first post!');
        $blogPost->setPublished(new \DateTime('2021-11-26 23:30:00'));
        $blogPost->setContent('Post text!');
        $blogPost->setAuthor($user);
        $blogPost->setSlug('a-first-post');

        $manager->persist($blogPost);

        $blogPost = new BlogPost();
        $blogPost->setTitle('A second post!');
        $blogPost->setPublished(new \DateTime('2021-11-26 23:31:00'));
        $blogPost->setContent('Post text!');
        $blogPost->setAuthor($user);
        $blogPost->setSlug('a-second-post');

        $manager->persist($blogPost);

        $manager->flush();
    }

    public function loadComments(ObjectManager $manager)
    {

    }

    public function loadUsers(ObjectManager $manager)
    {
        $user = new User();
        $user->setUsername('admin')
            ->setEmail('admin@blog.com')
            ->setName('Bruno Jesus');

        $user->setPassword($this->userPasswordHasher->hashPassword(
            $user,
            'secret123#'
        ));

        $this->addReference('user_admin', $user);

        $manager->persist($user);
        $manager->flush();
    }
}
```

> 💡 The `UserPasswordHasher` hashes password in **bcrypt**, the `User` entity has to implement the `PasswordAuthenticatedUserInterface`.
When creating the user we call the `addReference` method to later be able to use that object from the `loadBlogPosts` method.


**Execute the doctrine:fixtures:load command**:

```bash
php bin/console doctrine:fixtures:load
```

### Generating fake data

Instead of creating our fake data, we can generate it with a **faker** library.

**Install the ``faker`` library:**

```bash
composer install --dev fakerphp/faker
```

Fixture with faker:

```php
<?php

namespace App\DataFixtures;

use App\Entity\BlogPost;
use App\Entity\User;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;
use Faker\Factory;
use Faker\Generator;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;

class AppFixtures extends Fixture
{

    private UserPasswordHasherInterface $userPasswordHasher;
    private Generator $faker;

    public function __construct(UserPasswordHasherInterface $userPasswordHasher)
    {
        $this->userPasswordHasher = $userPasswordHasher;
        $this->faker = Factory::create();
    }

    public function load(ObjectManager $manager): void
    {
        $this->loadUsers($manager);
        $this->loadBlogPosts($manager);
    }

    public function loadBlogPosts(ObjectManager $manager)
    {
        /** @var User $user */
        $user = $this->getReference('user_admin');

        $numberOfPosts = $this->faker->numberBetween(50, 100);

        for ($i = 0; $i < $numberOfPosts; $i++) {
            $blogPost = new BlogPost();
            $blogPost->setTitle($this->faker->realText(30));
            $blogPost->setPublished($this->faker->dateTime());
            $blogPost->setContent($this->faker->realText());
            $blogPost->setAuthor($user);
            $blogPost->setSlug($this->faker->slug());

            $manager->persist($blogPost);
        }

        $manager->flush();
    }

    public function loadComments(ObjectManager $manager)
    {

    }

    public function loadUsers(ObjectManager $manager)
    {
        $user = new User();
        $user->setUsername('admin')
            ->setEmail('admin@blog.com')
            ->setName('Bruno Jesus');

        $user->setPassword($this->userPasswordHasher->hashPassword(
            $user,
            'secret123#'
        ));

        $this->addReference('user_admin', $user);

        $manager->persist($user);
        $manager->flush();
    }
}
```

## 🛼 EasyAdmin

EasyAdmin is a administration backoffice that can perform CRUD operations on your database.

### Installation
```bash
composer require easycorp/easyadmin-bundle
```

### Configuration

Generate the dashboard controller:
```bash
php bin/console make:admin:dashboard
```

Generate CRUD controllers:
```bash
php bin/console make:admin:crud
```

Edit the `DashboardController`:

```php
public function index(): Response
{
  // redirect to some CRUD controller
  $routeBuilder = $this->get(AdminUrlGenerator::class);
  return $this->redirect($routeBuilder->setController(BlogPostCrudController::class)->generateUrl());
}
```

## 🤯 API Platform

Framework for Creating API-driven projects. Uses **Symfony**.

### Installation
```bash
composer require api
```

### Configuration
Add the `@APIResource()` annotation to an entity that is a resource of the API.

### Usage

Open `http://localhost:8080/api` in your browser, you should see Swagger like documentation.

### ⛔ Restrict Operations

#### Collection Operations

**/api/users**

| Method | Description |
| --- | --- |
| GET | Get all elements (paginated) |
| POST | Create a new element |

**Restricting**

```php
<?php

use ApiPlatform\Core\Annotation\ApiResource;
...

/**
 * @ApiResource(
 *     itemOperations={"get", "post"},
 * )
 * @ORM\Entity(repositoryClass=UserRepository::class)
 */
class User implements PasswordAuthenticatedUserInterface
...
```
> 📌 Notice the `itemOperations={"get", "post"}` inside the `@ApiResource` annotation.

#### Item Operations

**/api/users/{id}**

| Method | Description |
| --- | --- |
| GET | Gets an element |
| PUT | Replaces an element |
| PATCH | Modifies an element |
| DELETE | Deletes an element |

**Restricting**

```php
<?php

use ApiPlatform\Core\Annotation\ApiResource;
...

/**
 * @ApiResource(
 *     itemOperations={"get", "post"},
 *     collectionOperations={"get", "post", "put", "patch", "delete"},
 * )
 * @ORM\Entity(repositoryClass=UserRepository::class)
 */
class User implements PasswordAuthenticatedUserInterface
...
```
> 📌 The restriction is being enforced by the `collectionOperations={"get", "post", "put", "patch", "delete"},` inside the `@ApiResource` annotation.

#### Restricting based on authentication

**Only allow authenticated users**

```php
/**
 * @ApiResource(
 *      itemOperations={
 *           "get"={
 *               "access_control"="is_granted('IS_AUTHENTICATED_FULLY')"
 *            }
 *       },
 *      collectionOperations={"post"},
 *      normalizationContext={
 *            "groups"={"read"}
 *      }
 *  )
 * @ORM\Entity(repositoryClass=UserRepository::class)
 * @UniqueEntity("username")
 * @UniqueEntity("email")
 * @method string getUserIdentifier()
*/
class User implements UserInterface, PasswordAuthenticatedUserInterface
{
```
> 📌 The authenticated only restriction is done by the `"access_control"="is_granted('IS_AUTHENTICATED_FULLY')"` inside the `itemOperations.get` parameter of the `@ApiResource` annotation.

**Only allow user responsible for the resource**

```php
/**
 * @ApiResource(
 *      itemOperations={
 *           "get",
 *           "put"={
 *               "security"="is_granted('IS_AUTHENTICATED_FULLY') and object.getAuthor() == user"
 *           }
 *      },
 *      collectionOperations={
 *           "get",
 *           "post"={
 *               "security"="is_granted('IS_AUTHENTICATED_FULLY')"
 *           }
 *       }
 *  )
 * @ORM\Entity(repositoryClass=BlogPostRepository::class)
 */
class BlogPost
{

...

	/**
     * @ORM\ManyToOne(targetEntity="App\Entity\User")
     * @ORM\JoinColumn(nullable=false)
     */
	private $author;
```

### 🤐 Serialization Groups
Sometimes we need to filter some sensitive fields, we can do that with groups:

```php
<?php

namespace App\Entity;

use ApiPlatform\Core\Annotation\ApiResource;
use App\Repository\UserRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Security\Core\User\PasswordAuthenticatedUserInterface;
use Symfony\Component\Serializer\Annotation\Groups;

/**
 * @ApiResource(
 *     itemOperations={"get"},
 *     collectionOperations={},
 *     normalizationContext={
 *           "groups"={"read"}
 *     }
 * )
 * @ORM\Entity(repositoryClass=UserRepository::class)
 */
class User implements PasswordAuthenticatedUserInterface
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     * @Groups({"read"})
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=255)
     * @Groups({"read"})
     */
    private $username;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $password;

    /**
     * @ORM\Column(type="string", length=255)
     * @Groups({"read"})
     */
    private $name;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $email;

    /**
     * @ORM\OneToMany(targetEntity="App\Entity\BlogPost", mappedBy="author")
     * @Groups({"read"})
     */
    private $posts;

    /**
     * @ORM\OneToMany(targetEntity="App\Entity\Comment", mappedBy="author")
     * @Groups({"read"})
     */
    private $comments;

    public function __construct()
    {
        $this->posts = new ArrayCollection();
        $this->comments = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getUsername(): ?string
    {
        return $this->username;
    }

    public function setUsername(string $username): self
    {
        $this->username = $username;

        return $this;
    }

    public function getPassword(): ?string
    {
        return $this->password;
    }

    public function setPassword(string $password): self
    {
        $this->password = $password;

        return $this;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(string $name): self
    {
        $this->name = $name;

        return $this;
    }

    public function getEmail(): ?string
    {
        return $this->email;
    }

    public function setEmail(string $email): self
    {
        $this->email = $email;

        return $this;
    }

    /**
     * @return Collection
     */
    public function getPosts(): Collection
    {
        return $this->posts;
    }

    /**
     * @return Collection
     */
    public function getComments(): Collection
    {
        return $this->comments;
    }

}
```
> 📌 The groups are defined inside the `normalizationContext` that you can find inside the `@ApiResource`. We then add all fields except `$email` and `$password` to the `read` group.

#### Using multiple groups, and include denormalization

```php
<?php

namespace App\Entity;

use ApiPlatform\Core\Annotation\ApiResource;
use App\Repository\UserRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Bridge\Doctrine\Validator\Constraints\UniqueEntity;
use Symfony\Component\Security\Core\User\PasswordAuthenticatedUserInterface;
use Symfony\Component\Security\Core\User\UserInterface;
use Symfony\Component\Serializer\Annotation\Groups;
use Symfony\Component\Validator\Constraints as Assert;

/**
 * @ApiResource(
 *      normalizationContext={"groups"={"get"}},
 *      itemOperations={
 *           "get"={
 *               "security"="is_granted('IS_AUTHENTICATED_FULLY')",
 *               "normalization_context"={
 *                   "groups"={"get"}
 *                }
 *            },
 *            "put"={
 *               "security"="is_granted('IS_AUTHENTICATED_FULLY') and object.getUsername() == user.getUsername()",
 *               "denormalization_context"={
 *                   "groups"={"put"}
 *                }
 *           }
 *       },
 *      collectionOperations={
 *           "post"={
 *                "denormalization_context"={
 *                   "groups"={"post"}
 *                }
 *           }
 *      },
 *  )
 * @ORM\Entity(repositoryClass=UserRepository::class)
 * @UniqueEntity("username")
 * @UniqueEntity("email")
 */
class User implements UserInterface, PasswordAuthenticatedUserInterface
{
    /**
     * @ORM\Id
	 * @ORM\GeneratedValue
	 * @ORM\Column(type="integer")
     * @Groups({"get"})
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=255)
     * @Groups({"get", "post"})
     * @Assert\NotBlank()
     * @Assert\Length(min=6, max=255)
     */
    private $username;

    /**
     * @ORM\Column(type="string", length=255)
     * @Groups({"put", "post"})
     * @Assert\NotBlank()
     * @Assert\Regex(
     *      pattern="/(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{7,}/",
     *      message="Password must be seven characters long and contains at least one digit, one upper case and one lower case letter"
     *  )
     */
    private $password;

    /**
     * @Groups({"put", "post"})
     * @Assert\NotBlank
     * @Assert\Expression(
     *      "this.getPassword() === this.getRetypedPassword()",
     *      message="Passwords does not match"
     *  )
     */
    private $retypedPassword;

    /**
     * @ORM\Column(type="string", length=255)
     * @Groups({"get", "post", "put"})
     * @Assert\NotBlank()
     * @Assert\Length(min=3, max=255)
     */
    private $name;

    /**
     * @ORM\Column(type="string", length=255)
     * @Groups({"post", "put"})
     * @Assert\NotBlank()
     * @Assert\Email()
     * @Assert\Length(min=6, max=255)
     */
    private $email;

    /**
     * @ORM\OneToMany(targetEntity="App\Entity\BlogPost", mappedBy="author")
     * @Groups({"get"})
     */
    private $posts;

    /**
     * @ORM\OneToMany(targetEntity="App\Entity\Comment", mappedBy="author")
     * @Groups({"get"})
     */
    private $comments;

...
```

### 🧒 Api Sub-Resource

Consider the **BlogPost**

```php
class BlogPost implements AuthoredEntityInterface, PublishedDateEntityInterface
{
/**
 * @ORM\Id
 * @ORM\GeneratedValue
 * @ORM\Column(type="integer")
 */
private $id;

/**
 * @ORM\Column(type="string", length=255)
 * @Assert\Length(min=10)
 * @Groups({"post"})
 */
private $title;

/**
 * @ORM\Column(type="datetime")
 */
private $published;

/**
 * @ORM\Column(type="text")
 * @Assert\NotBlank()
 * @Assert\Length(min=20)
 * @Groups({"post"})
 */
private $content;

/**
 * @ORM\ManyToOne(targetEntity="App\Entity\User")
 * @ORM\JoinColumn(nullable=false)
 */
private $author;

/**
 * @ORM\Column(type="string", length=255)
 * @Assert\NotBlank()
 * @Groups({"post"})
 */
private $slug;

/**
 * @ORM\OneToMany(targetEntity="App\Entity\Comment", mappedBy="blogPost")
 */
private Collection $comments;

...
```

When queried with **ApiPlaftorm** we will get something like this:

`GET http://localhost:8080/api/blog_posts/583`

```json
{
	"@context": "\/api\/contexts\/BlogPost",
	"@id": "\/api\/blog_posts\/583",
	"@type": "BlogPost",
	"id": 583,
	"title": "Rabbit began. Alice thought.",
	"published": "2021-12-09T23:20:46+00:00",
	"content": "Mock Turtle would be grand, certainly,' said Alice, looking down at her for a dunce? Go on!' 'I'm a poor man, your Majesty,' the Hatter hurriedly left the court, without even waiting to put it.",
	"author": "\/api\/users\/22",
	"slug": "nostrum-minus-aut-harum-autem-voluptas",
	"comments": [
		"\/api\/comments\/764",
		"\/api\/comments\/765",
		"\/api\/comments\/766",
		"\/api\/comments\/767",
		"\/api\/comments\/768",
		"\/api\/comments\/769",
		"\/api\/comments\/770",
		"\/api\/comments\/771"
	]
}
```

With this approach we are seeing the comment uris, but we have no way of getting all the comments for this **BlogPost** without getting them one by one.

To accomplish that we have the **@ApiSubresource** annotation:

```php
class BlogPost implements AuthoredEntityInterface, PublishedDateEntityInterface
{

...

/**
 * @ORM\OneToMany(targetEntity="App\Entity\Comment", mappedBy="blogPost")
 * @ApiSubresource()
 */
private Collection $comments;

...
```

We now have the following **route**:

`GET https://localhost:8080/api/blog_posts/583/comments`

```json
{
	"@context": "\/api\/contexts\/Comment",
	"@id": "\/api\/blog_posts\/583\/comments",
	"@type": "hydra:Collection",
	"hydra:member": [
		{
			"@id": "\/api\/comments\/764",
			"@type": "Comment",
			"id": 764,
			"content": "Do you think, at your age, it is I hate cats and dogs.' It was the Cat in a solemn tone, 'For the Duchess. An invitation from the roof. There were doors all round her, about the whiting!' 'Oh, as to.",
			"published": "2021-10-18T20:12:08+00:00",
			"author": "\/api\/users\/22",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/765",
			"@type": "Comment",
			"id": 765,
			"content": "Alice. 'I don't know the meaning of it had finished this short speech, they all moved off, and she very soon had to be executed for having missed their turns, and she went on, half to Alice. 'What.",
			"published": "2021-10-11T20:39:00+00:00",
			"author": "\/api\/users\/22",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/766",
			"@type": "Comment",
			"id": 766,
			"content": "Alice heard the King said to herself, for she felt certain it must be collected at once took up the little passage: and THEN--she found herself lying on the trumpet, and then Alice put down yet.",
			"published": "2021-07-25T16:48:45+00:00",
			"author": "\/api\/users\/21",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/767",
			"@type": "Comment",
			"id": 767,
			"content": "Dormouse, after thinking a minute or two, and the cool fountains. CHAPTER VIII. The Queen's argument was, that you couldn't cut off a bit of the other side of the deepest contempt. 'I've seen a good.",
			"published": "2021-08-19T05:12:46+00:00",
			"author": "\/api\/users\/22",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/768",
			"@type": "Comment",
			"id": 768,
			"content": "So she went on for some time in silence: at last came a little sharp bark just over her head pressing against the door, she walked off, leaving Alice alone with the lobsters to the table to measure.",
			"published": "2021-07-16T09:12:05+00:00",
			"author": "\/api\/users\/23",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/769",
			"@type": "Comment",
			"id": 769,
			"content": "Crab, a little quicker. 'What a number of executions the Queen had never forgotten that, if you wouldn't keep appearing and vanishing so suddenly: you make one repeat lessons!' thought Alice; 'I.",
			"published": "2021-12-20T20:29:51+00:00",
			"author": "\/api\/users\/21",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/770",
			"@type": "Comment",
			"id": 770,
			"content": "Alice heard it muttering to himself in an impatient tone: 'explanations take such a fall as this, I shall fall right THROUGH the earth! How funny it'll seem to put it to her ear. 'You're thinking.",
			"published": "2021-12-27T07:13:00+00:00",
			"author": "\/api\/users\/21",
			"blogPost": "\/api\/blog_posts\/583"
		},
		{
			"@id": "\/api\/comments\/771",
			"@type": "Comment",
			"id": 771,
			"content": "The Caterpillar was the same age as herself, to see if she meant to take the place of the miserable Mock Turtle. 'Very much indeed,' said Alice. 'Why, there they lay on the breeze that followed.",
			"published": "2021-05-23T11:10:30+00:00",
			"author": "\/api\/users\/24",
			"blogPost": "\/api\/blog_posts\/583"
		}
	],
	"hydra:totalItems": 8
}
```

#### Displaying nested Sub-Resources

Sometimes we may need to display some information of a subresource, like this:

```json
{
	"@context": "\/api\/contexts\/Comment",
	"@id": "\/api\/blog_posts\/583\/comments",
	"@type": "hydra:Collection",
	"hydra:member": [
		{
			"@id": "\/api\/comments\/764",
			"@type": "Comment",
			"content": "Do you think, at your age, it is I hate cats and dogs.' It was the Cat in a solemn tone, 'For the Duchess. An invitation from the roof. There were doors all round her, about the whiting!' 'Oh, as to.",
			"published": "2021-10-18T20:12:08+00:00",
			"author": {
				"@id": "\/api\/users\/22",
				"@type": "User",
				"username": "john_doe",
				"name": "John Doe"
			}
		}
	],
	"hydra:totalItems": 1
}
```
> 📌 `author` is a subresource of `comment`, you can see the full `author` in the `json` response.

We can accomplish that by using **subresourceOperations**.

##### How to do it

###### For a specific route (optional)

First we need to get the route on which we want to have the attributes of the sub-resource:

```bash
php bin/console debug:route
```

```bash
----------------------------------------- -------- -------- ------ ----------------------------------------- 
 Name                                      Method   Scheme   Host   Path                                     
----------------------------------------- -------- -------- ------ ----------------------------------------- 
 ...
 api_blog_posts_comments_get_subresource   GET      ANY      ANY    /api/blog_posts/{id}/comments.{_format} 
 ...
----------------------------------------- -------- -------- ------ -----------------------------------------
```

Then we define the subresource operation in our **ApiResource:**

>❗ If we do not want this for a spefic route, there’s no need to use the **subresourceOperations**, just assign the **normalization_context** to an **itemOperations**.


```php
/**
 * @ApiResource(
 *     itemOperations={
 *          "get",
 *          "put"={
 *              "security"="is_granted('IS_AUTHENTICATED_FULLY') and object.getAuthor() == user"
 *          }
 *     },
 *     collectionOperations={
 *          "get",
 *          "post"={
 *              "security"="is_granted('IS_AUTHENTICATED_FULLY')"
 *          }
 *     },
 *     subresourceOperations={
 *          "api_blog_posts_comments_get_subresource"={
 *              "normalization_context"={
 *                  "groups"={"get-comment-with-author"}
 *              }
 *          }
 *     },
 *     denormalizationContext={
 *      "groups"={"post"}
 *     }
 * )
 * @ORM\Entity(repositoryClass=CommentRepository::class)
 */
class Comment implements AuthoredEntityInterface, PublishedDateEntityInterface
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     * @Groups({"get-comment-with-authorget-comment-with-author"})
     */
    private $id;

    /**
     * @ORM\Column(type="text")
     * @Assert\NotBlank()
     * @Assert\Length(min=5, max=3000)
     * @Groups({"get-comment-with-author","post"})
     */
    private $content;

    /**
     * @ORM\Column(type="datetime")
     * @Groups({"get-comment-with-author"})
     */
    private $published;

    /**
     * @ORM\ManyToOne(targetEntity="App\Entity\User", inversedBy="comments")
     * @ORM\JoinColumn(nullable=false)
     * @Groups({"get-comment-with-author"})
     */
    private $author;

    /**
     * @ORM\ManyToOne(targetEntity="App\Entity\BlogPost")
     * @ORM\JoinColumn(nullable=false)
     * @Groups({"post"})
     */
    private $blogPost;
```

We now just need to use the same group on our sub resources (`get-comment-with-authorget-comment-with-author` and `get-comment-with-author`):

```php
/**
 * @ApiResource(
 *      normalizationContext={"groups"={"get"}},
 *      itemOperations={
 *           "get"={
 *               "security"="is_granted('IS_AUTHENTICATED_FULLY')",
 *               "normalization_context"={
 *                   "groups"={"get"}
 *                }
 *            },
 *            "put"={
 *               "security"="is_granted('IS_AUTHENTICATED_FULLY') and object.getUsername() == user.getUsername()",
 *               "denormalization_context"={
 *                   "groups"={"put"}
 *                }
 *           }
 *       },
 *      collectionOperations={
 *           "post"={
 *                "denormalization_context"={
 *                   "groups"={"post"}
 *                }
 *           }
 *      },
 *  )
 * @ORM\Entity(repositoryClass=UserRepository::class)
 * @UniqueEntity("username")
 * @UniqueEntity("email")
 */
class User implements UserInterface, PasswordAuthenticatedUserInterface
{
	/**
     * @ORM\Id
	 * @ORM\GeneratedValue
	 * @ORM\Column(type="integer")
     * @Groups({"get"})
     */
	 private $id;

	/**
     * @ORM\Column(type="string", length=255)
     * @Groups({"get", "post", "get-comment-with-author"})
     * @Assert\NotBlank()
     * @Assert\Length(min=6, max=255)
     */
	private $username;

	/**
     * @ORM\Column(type="string", length=255)
     * @Groups({"put", "post"})
     * @Assert\NotBlank()
     * @Assert\Regex(
     *      pattern="/(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{7,}/",
     *      message="Password must be seven characters long and contains at least one digit, one upper case and one lower case letter"
     * )
     */
	 private $password;

	/**
     * @Groups({"put", "post"})
     * @Assert\NotBlank
	 * @Assert\Expression(
     *      "this.getPassword() === this.getRetypedPassword()",
     *      message="Passwords does not match"
     * )
     */
	 private $retypedPassword;

	/**
     * @ORM\Column(type="string", length=255)
     * @Groups({"get", "post", "put", "get-comment-with-author"})
     * @Assert\NotBlank()
     * @Assert\Length(min=3, max=255)
     */
	 private $name;

	/**
     * @ORM\Column(type="string", length=255)
     * @Groups({"post", "put"})
     * @Assert\NotBlank()
     * @Assert\Email()
     * @Assert\Length(min=6, max=255)
     */
	 private $email;

	/**
     * @ORM\OneToMany(targetEntity="App\Entity\BlogPost", mappedBy="author")
     * @Groups({"get"})
     */
	 private $posts;

	/**
     * @ORM\OneToMany(targetEntity="App\Entity\Comment", mappedBy="author")
     * @Groups({"get"})
     */
	 private $comments;
```
